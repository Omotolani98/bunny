package impl

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"

	"github.com/Omotolani98/bunny/internal/errors"
)

const BUNNYHOME = ".bunny"
type Hetzner struct {
	Token string `json:"hetzner_token"`
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
