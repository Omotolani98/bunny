package cmd

import (
	"github.com/Omotolani98/bunny/internal/cloud"
	"github.com/spf13/cobra"
)

func Init() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "bunny",
		Short: "Bunny is a cli tool that deploys application easily to your hetzner cloud vm and exposes port",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return cloud.InitBunnyDir()
		},
	}

	rootCmd.AddCommand(AuthCmd())
	return rootCmd
}
