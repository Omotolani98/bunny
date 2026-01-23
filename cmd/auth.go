package cmd

import (
	"fmt"
	"strings"

	"github.com/Omotolani98/bunny/internal/cloud"
	"github.com/Omotolani98/bunny/internal/cloud/impl"
	"github.com/spf13/cobra"
)

const AuthMsg = `
	Bunny Auth

	Here's a simple guide to get your token - %s

	We need you to input your %s Token here: `

var (
	cloudProvider string
)

const (
	hetznerLink = "https://docs.hetzner.com/cloud/api/getting-started/generating-api-token/"
	awsLink = "https://aws.amazon.co.uk/guides"
)

func AuthCmd() *cobra.Command {
	authCmd := &cobra.Command{
		Use: "auth",
		Short: "Authenticate with cloud provider token",
		RunE: func(cmd *cobra.Command, args []string) error {
			cloudProvider = strings.ToLower(cloudProvider)

			var input string
			fmt.Printf(AuthMsg, GetProviderGuideLink(cloudProvider), cloudProvider)
			fmt.Scan(&input)

			h := impl.Hetzner{
				Token: input,
			}

			cm := cloud.New(h)
			res, err := cm.Auth()
			fmt.Print(res)
			return err
		},
	}

	authCmd.Flags().StringVarP(&cloudProvider, "cloud", "c", "hetzner", "specify cloud provider")
	return authCmd
}

func GetProviderGuideLink(p string) string {
	switch p {
	case "aws":
		return awsLink
	case "azure":
		return ""
	default:
		return hetznerLink
	}
}
