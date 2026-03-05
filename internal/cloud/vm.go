package cloud

import "golang.org/x/crypto/ssh"

type VMConnection struct {
	Host string
	User string
	Key  string
}

func (v *VMConnection) SSHConnect() (*ssh.Client, error) {
	// SSH into the VM
	return nil, nil
}

func (v *VMConnection) Deploy(image string, port string, env map[string]string) error {
	// pull image, run container
	return nil
}
