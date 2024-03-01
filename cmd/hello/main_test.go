package main

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestRun(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	set.Bool("verbose", false, "doc")
	set.Uint64("exclnum", 1, "doc")
	c := cli.NewContext(nil, set, nil)
	err := set.Parse([]string{"World"})
	assert.Nil(t, err)

	err = Run(c)
	assert.Nil(t, err)
}
