package main

import (
	"os"

	"github.com/npalladium/ccc/internal/commands"
	_ "github.com/npalladium/ccc/internal/config"
)

var version = "main"

func main() {
	os.Exit(commands.Run(version))
}
