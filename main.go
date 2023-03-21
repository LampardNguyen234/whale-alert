package main

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/cmd"
	"os"
)

func main() {
	if err := cmd.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
