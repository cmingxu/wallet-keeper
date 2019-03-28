package main

import (
	"github.com/urfave/cli"
)

var WebCmd cli.Command = cli.Command{
	Name:    "web",
	Aliases: []string{"w"},
	Usage:   "run jex web application",
	Flags: []cli.Flag{
		httpAddrFlag,
		logLevelFlag,
		logPathFlag,
	},
	Action: func(c *cli.Context) error {
		log.Println(c.StringFlag("logLevel"))
		return nil
	},
}
