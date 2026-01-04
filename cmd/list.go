package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/dalryan/ip-enrich/internal/provider"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available intelligence providers",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		providers := provider.All()
		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 3, ' ', 0)
		if _, err := fmt.Fprintln(w, "ID\tNAME"); err != nil {
			return err
		}
		if _, err := fmt.Fprintln(w, "--\t----"); err != nil {
			return err
		}
		for _, p := range providers {
			if _, err := fmt.Fprintf(w, "%s\t%s\n", p.ID(), p.Name()); err != nil {
				return err
			}
		}
		return w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
