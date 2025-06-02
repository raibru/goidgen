package main

import (
	"os"

	"github.com/raibru/goidgen/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
