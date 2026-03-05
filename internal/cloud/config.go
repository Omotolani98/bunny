package cloud

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func loadConfig() *BunnyConfig {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".bunny", "config.json")

	data, err := os.ReadFile(path)
	if err != nil {
		return &BunnyConfig{
			Version: "1",
			Clouds:  make(map[string]CloudProp),
			Apps:    make(map[string]AppProp),
			VMs:     make(map[string]VMProp),
		}
	}

	var config BunnyConfig
	json.Unmarshal(data, &config)
	return &config
}

func InitBunnyDir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dir := filepath.Join(home, ".bunny")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	configPath := filepath.Join(dir, "config.json")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		config := BunnyConfig{
			Version: "1",
			Clouds:  make(map[string]CloudProp),
			Apps:    make(map[string]AppProp),
			VMs:     make(map[string]VMProp),
		}

		data, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return err
		}

		return os.WriteFile(configPath, data, 0644)
	}

	return nil
}
