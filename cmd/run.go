package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/loomnetwork/e2e/engine"
	"github.com/loomnetwork/e2e/lib"
	"github.com/loomnetwork/e2e/node"
	"github.com/spf13/cobra"
)

func newRunCommand() *cobra.Command {
	var confFilename string
	command := &cobra.Command{
		Use:           "run",
		Short:         "Run nodes",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := lib.ReadConfig(confFilename)
			if err != nil {
				return err
			}

			// Trap Interrupts, SIGINTs and SIGTERMs.
			sigC := make(chan os.Signal, 1)
			signal.Notify(sigC, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			defer signal.Stop(sigC)

			errC := make(chan error)
			eventC := make(chan *node.Event)
			e := engine.New(conf)

			ctx, cancel := context.WithCancel(context.Background())
			go func() { errC <- e.Run(ctx, eventC) }()

			func() {
				for {
					select {
					case err := <-errC:
						fmt.Printf("error: %s\n", err)
						cancel()
						return
					case <-sigC:
						cancel()
						fmt.Printf("stopping runner\n")
						return
					}
				}
			}()

			return nil
		},
	}

	flags := command.Flags()
	flags.StringVar(&confFilename, "conf", "default/runner.toml", "Runner configuration path")
	return command
}
