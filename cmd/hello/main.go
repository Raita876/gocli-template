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

type Arguments struct {
	Message string
}

type Argument interface {
	apply(*Arguments)
}

type messageArg string

func (ma messageArg) apply(a *Arguments) {
	a.Message = string(ma)
}

func (arguments *Arguments) Set(args ...Argument) {
	for _, a := range args {
		a.apply(arguments)
	}
}

type Options struct {
	Verbose bool
	ExclNum uint8
}

type Option interface {
	apply(*Options)
}

type verboseOption bool

func (vo verboseOption) apply(o *Options) {
	o.Verbose = bool(vo)
}

type exclNumOption uint8

func (eo exclNumOption) apply(o *Options) {
	o.ExclNum = uint8(eo)
}

func (options *Options) Set(opts ...Option) {
	for _, o := range opts {
		o.apply(options)
	}
}

func Hello(c *cli.Context) error {
	a := &Arguments{}
	o := &Options{}
	a.Set(
		messageArg(c.Args().Get(0)),
	)
	o.Set(
		verboseOption(c.Bool("verbose")),
		exclNumOption(c.Uint64("exclnum")),
	)

	return hello.Hello()
}

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Version: version,
		Name:    name,
		Usage:   "This is Hello CLI.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Value:   false,
				Usage:   "Output detailed log.",
			},
			&cli.Uint64Flag{
				Name:    "exclnum",
				Aliases: []string{"n"},
				Value:   1,
				Usage:   "Output detailed log.",
			},
		},
		Action: Hello,
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
