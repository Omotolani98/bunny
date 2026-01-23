package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Omotolani98/bunny/internal/cloud/impl"
	"github.com/spf13/cobra"
)

func LocationCmd() *cobra.Command {
	locationCmd := &cobra.Command{
		Use: "location",
		Short: "Displays list of available locations",
		RunE: func(cmd *cobra.Command, args []string) error {	
			h := impl.NewHetzner()
			res, err := h.Locations()
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

			fmt.Fprintln(w, "NAME\tCITY\tCOUNTRY\tZONE\tDESCRIPTION")
			fmt.Fprintln(w, "----\t-------------\t-------\t------------\t---------------------")

			for _, v := range res {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
					v.Name,
					v.City,
					v.Country,
					v.NetworkZone,
					v.Description,
				)
			}

			w.Flush()
			return err
		},
	}

	return locationCmd
}
