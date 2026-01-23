package cmd

import (
	"fmt"

	"github.com/Omotolani98/bunny/internal/cloud/impl"
	"github.com/spf13/cobra"
)

const AuthMsg = `
	Bunny Auth

	Here's a simple guide to get your token - https://docs.hetzner.com/cloud/api/getting-started/generating-api-token/

	We need you to input your Hetzner Token here: `

// var (
// 	cloudProvider string
// )

func AuthCmd() *cobra.Command {
	authCmd := &cobra.Command{
		Use: "auth",
		Short: "Authenticate with cloud Hetzner token",
		RunE: func(cmd *cobra.Command, args []string) error {
			// cloudProvider = strings.ToLower(cloudProvider)

			var input string
			fmt.Printf(AuthMsg)
			fmt.Scan(&input)

			h := impl.NewHetzner()

			res, err := h.Auth()
			fmt.Print(res)
			return err
		},
	}

	// authCmd.Flags().StringVarP(&cloudProvider, "cloud", "c", "hetzner", "specify cloud provider")
	return authCmd
}

