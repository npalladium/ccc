package main

import (
	"os"

	"github.com/npalladium/ccc/internal/commands"
)

var version = "main"

func main() {
	os.Exit(commands.Run(version))
}
