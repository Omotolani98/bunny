package cmd

import (
	"fmt"
	"log"

	"github.com/Omotolani98/bunny/internal/cloud"
	"github.com/Omotolani98/bunny/internal/tui"
	"github.com/spf13/cobra"
)

const AuthMsg = `
	Bunny Auth

	Here's a simple guide to get your token - https://docs.hetzner.com/cloud/api/getting-started/generating-api-token/

	We need you to input your Hetzner Token here: `

var (
	cloudProvider string
	host          string
	user          string
	key           string
	name          string
)

func AuthCmd() *cobra.Command {
	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate with cloud Hetzner token",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cloudProvider != "" {
				credentials, err := tui.PromptCloudAuth(cloudProvider)
				if err != nil {
					log.Fatal(err)
				}

				// region := credentials["region"] // empty for hetzner, that's fine
				manager := cloud.NewCloudManagerWithProvider(cloudProvider, credentials)
				manager.SaveCloudAuth(cloudProvider, credentials)

				fmt.Printf("✓ Authenticated with %s\n", cloudProvider)
				return nil
			} else {
				manager := cloud.NewCloudManagerWithVM(host, user, key)
				manager.RegisterVM(name, host, user, key)
				return nil
			}
		},
	}

	authCmd.Flags().StringVarP(&cloudProvider, "cloud", "c", "", "specify cloud provider")
	authCmd.Flags().StringVarP(&host, "host", "H", "", "ip of existing vm")
	authCmd.Flags().StringVarP(&name, "name", "n", "", "name of vm")
	authCmd.Flags().StringVarP(&user, "user", "u", "", "user ssh")
	authCmd.Flags().StringVarP(&key, "key", "k", "", "ssh key to use to access vm")
	return authCmd
}

func promptForToken() {

}
