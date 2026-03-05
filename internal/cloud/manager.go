package cloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Omotolani98/bunny/internal/provider/hetzner"
)

type CloudManager struct {
	Config   *BunnyConfig
	Provider Cloud
}

func NewCloudManagerWithVM(host, user, key string) *CloudManager {
	return &CloudManager{
		Config:   loadConfig(),
		Provider: nil,
	}
}

func NewCloudManagerWithProvider(providerName string, credentials map[string]string) *CloudManager {
	var provider Cloud

	switch providerName {
	case "hetzner":
		provider = hetzner.NewHetzner(credentials["token"])
	}

	return &CloudManager{
		Config:   loadConfig(),
		Provider: provider,
	}
}

func GetCloud(cloudProvider string) (Cloud, error) {
	switch cloudProvider {
	case "hetzner":
		return hetzner.Hetzner{}, nil
	default:
		return nil, errors.New("cloud provider not recognized")
	}
}

func (m *CloudManager) RegisterVM(name, host, user, key string) error {
	if m.Config.VMs == nil {
		m.Config.VMs = make(map[string]VMProp)
	}

	m.Config.VMs[name] = VMProp{
		Host:         host,
		User:         user,
		Key:          key,
		RegisteredAt: time.Now().UTC().Format(time.RFC3339),
	}

	return m.SaveConfig()
}

func (m *CloudManager) RemoveVM(name string, removeApps bool) error {
	if m.Config.VMs == nil {
		return fmt.Errorf("no VMs registered")
	}

	if _, exists := m.Config.VMs[name]; !exists {
		return fmt.Errorf("VM '%s' not found", name)
	}

	if removeApps {
		for appName, app := range m.Config.Apps {
			if app.VM == name {
				delete(m.Config.Apps, appName)
			}
		}
	}

	delete(m.Config.VMs, name)

	return m.SaveConfig()
}

func (m *CloudManager) SaveCloudAuth(cloudProvider string, credentials map[string]string) error {
	if m.Config.Clouds == nil {
		m.Config.Clouds = make(map[string]CloudProp)
	}

	prop := CloudProp{
		AuthenticatedAt: time.Now().UTC().Format(time.RFC3339),
	}

	switch cloudProvider {
	case "hetzner":
		prop.Token = credentials["token"]
	case "aws":
		prop.AccessKeyID = credentials["access_key_id"]
		prop.SecretAccessKey = credentials["secret_access_key"]
		prop.Region = credentials["region"]
	case "gcp":
		prop.ServiceAccountJSON = credentials["service_account_json"]
	case "azure":
		prop.TenantID = credentials["tenant_id"]
		prop.ClientID = credentials["client_id"]
		prop.ClientSecret = credentials["client_secret"]
	default:
		return fmt.Errorf("unsupported cloud provider: %s", cloudProvider)
	}

	m.Config.Clouds[cloudProvider] = prop

	return m.SaveConfig()
}

func (m *CloudManager) SaveConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dir := filepath.Join(home, ".bunny")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(m.Config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(dir, "config.json"), data, 0644)
}
