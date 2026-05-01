package main

import (
	"os"
	"github.com/aleks-ship-it/pave/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
