package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "obs",
	Short:   "obs - CLI to open, search, move, create and update notes",
	Version: "v0.1.6",
	Long:    "obs - CLI to open, search, move, create and update notes",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		if err != nil {
			os.Exit(1)
		}
	}
}
