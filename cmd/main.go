package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "e2e-test",
		Short: "e2e-test utility",
	}

	rootCmd.AddCommand(newNewCommand())
	rootCmd.AddCommand(newRunCommand())
	rootCmd.AddCommand(newTestCommand())
	rootCmd.AddCommand(newGenerateCommand())
	rootCmd.AddCommand(newPubKeyCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
