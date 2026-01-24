package impl

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"slices"

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

func (h *Hetzner) Types(location string) ([]ServerType, error) {
	client, err := InitClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.ServerType.All(context.Background())
	if err != nil {
		return nil, err
	}

	// serverTypes := make([]ServerType, 0)
	// for _, v := range resp {
	// 	st := ServerType{
	// 		Type: v.Name,
	// 		Cpu: fmt.Sprintf("%d vCPU", v.Cores),
	// 		Ram: fmt.Sprintf("%f GB", v.Memory),
	// 		Disk: fmt.Sprintf("%d GB", v.Disk),
	// 		Price: v.Pricings[0].Monthly.Net,
	// 	}
	// 	serverTypes = append(serverTypes, st)
	// }
	//
	serverTypes := make([]ServerType, 0)

	for _, v := range resp {	
		var price string

		if location != "" {
			found := false
			for _, loc := range v.Locations {
				if loc.Location.Name == location {
					found = true
					break
				}
			}
			if !found {
				continue
			}

			for _, p := range v.Pricings {
				if p.Location != nil && p.Location.Name == location {
					price = p.Monthly.Net
					break
				}
			}
		} else {
			if len(v.Pricings) > 0 {
				price = v.Pricings[0].Monthly.Net
			}
		}

		st := ServerType{
			Type:  v.Name,
			Cpu:   fmt.Sprintf("%d vCPU", v.Cores),
			Ram:   fmt.Sprintf("%.1f GB", v.Memory),
			Disk:  fmt.Sprintf("%d GB", v.Disk),
			Price: price,
		}
		serverTypes = append(serverTypes, st)
	}

	slices.SortFunc(serverTypes, func(a, b ServerType) int {
		return cmp.Compare(a.Type, b.Type)
	})

	return serverTypes, nil
}
