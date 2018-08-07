package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() error {
	rootCmd := &cobra.Command{
		Use:   "e2e-test",
		Short: "e2e-test utility",
	}

	rootCmd.AddCommand(newNewCommand())
	rootCmd.AddCommand(newRunCommand())
	rootCmd.AddCommand(newTestCommand())
	rootCmd.AddCommand(newGenerateCommand())
	rootCmd.AddCommand(newPubKeyCommand())

	return rootCmd.Execute()
}
