package cmd

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dalryan/ip-enrich/internal/sources"
	"github.com/dalryan/ip-enrich/internal/ui"
	"github.com/dalryan/ip-enrich/internal/utils"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var ip string

var rootCmd = &cobra.Command{
	Use:   "ip-enrich [ip]",
	Short: "A quick and dirty tool for enriching an IP Address",
	Args:  cobra.MaximumNArgs(1),
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		ip = args[0]
	}
	if err := utils.ValidateIPAddr(ip); err != nil {
		fmt.Printf("Not a valid IP address: %s", ip)
		return
	}

	initialModel := ui.Model{
		Ip:              ip,
		ViewingResponse: false,
		Results:         make(map[string]sources.Result),
		Spinner:         spinner.New(spinner.WithSpinner(spinner.Jump)),
	}

	for i, endpoint := range ui.Endpoints {
		updatedURL := strings.Replace(endpoint.URL, "{ip}", ip, -1)
		ui.Endpoints[i].URL = updatedURL
		initialModel.Results[endpoint.Name] = sources.Result{
			Data: nil,
			Url:  updatedURL,
			Code: 0,
			Done: false,
		}
	}

	initialModel.Spinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	if _, err := tea.NewProgram(initialModel, tea.WithAltScreen(), tea.WithMouseCellMotion()).Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "", "IP address to enrich")
}
