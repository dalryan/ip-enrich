package main

import (
	"os"

	"github.com/dalryan/ip-enrich/cmd"
)

// main is the entrypoint for the tool
func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
