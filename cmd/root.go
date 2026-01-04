package cmd

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dalryan/ip-enrich/internal/output"
	"github.com/dalryan/ip-enrich/internal/provider"
	_ "github.com/dalryan/ip-enrich/internal/providers"
	"github.com/spf13/cobra"
)

var (
	outputFormat   string
	providerFilter []string
	timeout        int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ip-enrich [ip]",
	Short: "Aggregates threat intelligence for an IP address",
	Long: `A high-performance aggregator for IP threat intelligence.
    
Examples:
  ip-enrich 1.1.1.1
  ip-enrich 1.1.1.1 -p shodan,greynoise
  ip-enrich -o json 8.8.8.8`,

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		ip := args[0]

		if net.ParseIP(ip) == nil {
			return fmt.Errorf("'%s' is not a valid IP address", ip)
		}

		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		go func() {
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
			<-sigChan
			cancel()
		}()

		if len(providerFilter) > 0 {
			unknown := provider.Validate(providerFilter)
			if len(unknown) > 0 {
				cmd.PrintErrf("Warning: unknown providers ignored: %v\n", unknown)
			}
		}

		return run(ctx, ip, providerFilter, outputFormat, timeout, cmd.OutOrStdout())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

// init initialises all flags
func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&providerFilter, "providers", "p", []string{}, "Comma-separated list of providers")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "pretty", "Output format: json, pretty")
	rootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 10, "HTTP timeout in seconds")
}

// run takes the list of providers and executes them
func run(ctx context.Context, ip string, providerIDs []string, format string, timeoutSeconds int, w io.Writer) error {
	providers := provider.Filter(providerIDs)
	if len(providers) == 0 {
		return fmt.Errorf("no providers matched request")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	executor := provider.NewExecutor()
	results := executor.Execute(ctx, ip, providers, nil)

	report := output.NewReport(ip, time.Now().UTC().Format(time.RFC3339), results)

	formatter, err := output.GetFormatter(format, w)
	if err != nil {
		return err
	}

	return formatter.Format(report)
}
