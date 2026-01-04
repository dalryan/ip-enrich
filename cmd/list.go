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
	Long:  `Displays a table of all registered providers.`,

	Args: cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		providers := provider.All()
		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME")
		fmt.Fprintln(w, "--\t----")
		for _, p := range providers {
			fmt.Fprintf(w, "%s\t%s\n", p.ID(), p.Name())
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
