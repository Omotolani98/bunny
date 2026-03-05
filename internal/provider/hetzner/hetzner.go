package hetzner

import (
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

var client *hcloud.Client

const BUNNYHOME = ".bunny"

type Hetzner struct {
	Token string `json:"hetzner_token"`
}

func NewHetzner(token string) *Hetzner {
	return &Hetzner{
		Token: token,
	}
}

func (h Hetzner) Auth(token string) (string, error) {
	return "", nil
}
