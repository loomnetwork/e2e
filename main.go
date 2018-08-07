package main

import (
	"fmt"
	"os"

	"github.com/loomnetwork/e2e/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
