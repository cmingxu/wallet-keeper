package main

import (
	"github.com/urfave/cli"
)

var jobCmd cli.Command = cli.Command{
	Name:    "job",
	Aliases: []string{"j"},
	Usage:   "run jex job application",
	Flags: []cli.Flag{
		logLevelFlag,
		logPathFlag,
	},
	Action: func(c *cli.Context) error {

		return nil
	},
}
