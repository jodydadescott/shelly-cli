package main

import (
	"os"

	"github.com/jodydadescott/shelly-manager/shelly"
)

func main() {
	err := shelly.NewCmd().Execute()

	if err != nil {
		os.Exit(1)
	}
}
