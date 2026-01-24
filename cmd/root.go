package cmd

import (
	"github.com/spf13/cobra"
)

func Init() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "bunny",
		Short: "Bunny is a cli tool that deploys application easily to your hetzner cloud vm and exposes port",
	}

	rootCmd.AddCommand(AuthCmd())
	rootCmd.AddCommand(LocationCmd())
	rootCmd.AddCommand(TypesCmd())
	return rootCmd
}
