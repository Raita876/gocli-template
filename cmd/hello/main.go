package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/Raita876/hello/internal/hello"
)

var (
	version string
	name    string
)

func main() {
	app := &cli.App{
		Version: version,
		Name:    name,
		Usage:   "This is Hello CLI.",
		Action: func(c *cli.Context) error {
			return hello.Hello()
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
