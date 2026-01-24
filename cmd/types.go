package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/Omotolani98/bunny/internal/cloud/impl"
	"github.com/spf13/cobra"
)

var location string

func TypesCmd() *cobra.Command {
	typesCmd := &cobra.Command{
		Use: "types",
		Short: "VM Types based on different locations",
		RunE: func(cmd *cobra.Command, args []string) error {
			location = strings.ToLower(location)
			h := impl.NewHetzner()
			serverTypes, err := h.Types(location)

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

			fmt.Fprintln(w, "TYPE\tCPU\tRAM\tDISK\tPRICE")
			fmt.Fprintln(w, "----\t-------------\t-------\t------------\t---------------------")

			for _, v := range serverTypes {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
					v.Type,
					v.Cpu,
					v.Ram,
					v.Disk,
					v.Price,
				)
			}

			w.Flush()

			return err
		},
	}

	typesCmd.Flags().StringVarP(&location, "location", "l", "fsn1", "VM locations")
	return typesCmd
}
