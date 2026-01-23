package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/user"

	"github.com/Omotolani98/bunny/internal/errors"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

var client *hcloud.Client

const BUNNYHOME = ".bunny"
type Hetzner struct {
	Token string `json:"hetzner_token"`
}

func NewHetzner() *Hetzner {
	var h Hetzner
	home, err := getHomeDir()
	if err != nil {
		return nil 
	}
	bunnyPath := fmt.Sprintf("%s/%s", home, BUNNYHOME)
	credPath := fmt.Sprintf("%s/credentials.json", bunnyPath)

	b, err := os.ReadFile(credPath)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(b, &h)
	if err != nil {
		return nil
	}

	return &h
}

func InitClient() (*hcloud.Client, error) {	
	h := NewHetzner()

	client = hcloud.NewClient(
		hcloud.WithToken(h.Token),
	)

	return client, nil
}

func getHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	homeDir := usr.HomeDir
	fmt.Printf("HomeDir: %s\n", homeDir)
	return homeDir, nil
}

func (h Hetzner) Auth() (string, error) {
	if h.Token == "" {
		return "", errors.ErrEmptyToken
	}

	b, err := json.Marshal(h)
	if err != nil {
		return "", err
	}

	home, err := getHomeDir()
	if err != nil {
		return "", err
	}

	bunnyPath := fmt.Sprintf("%s/%s", home, BUNNYHOME)
	credPath := fmt.Sprintf("%s/credentials.json", bunnyPath)
	fmt.Printf("")
	
	err = os.WriteFile(credPath, b, 0644)
	if err != nil {
		return "", err
	}

	return "Auth Successful Twin!", nil
}

func (h *Hetzner) Locations() ([]*hcloud.Location, error) {
	client, err := InitClient()
	if err != nil {
		return nil, err
	}

	return client.Location.All(context.Background())
}
