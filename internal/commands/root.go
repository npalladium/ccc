package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Run(version string) int {
	app := &cli.App{
		Name:        "ccc",
		Version:     version,
		Description: "Cloud Cost Controller",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug, d",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:  "config",
				Usage: "Path to config file",
			},
		},
		CommandNotFound: cmdNotFound,
	}
	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		return 1
	}
	return 0
}

func cmdNotFound(c *cli.Context, command string) {
	fmt.Printf(
		"%s: '%s' is not a %s command. See '%s --help'.\n",
		c.App.Name,
		command,
		c.App.Name,
		os.Args[0],
	)
	os.Exit(1)
}
