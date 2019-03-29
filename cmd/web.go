package main

import (
	"strings"

	"github.com/urfave/cli"
)

var webCmd cli.Command = cli.Command{
	Name:    "web",
	Aliases: []string{"w"},
	Usage:   "run jex web application",
	Flags: []cli.Flag{
		httpAddrFlag,
		logLevelFlag,
		logPathFlag,
	},
	Action: func(c *cli.Context) error {
		var env string = strings.ToUpper(c.String("env"))

		// setup logger
		loggerFormat := "json"
		if env == "dev" {
			loggerFormat = "text"
		}
		err := setupLogger(c.String("log-level"), c.String("log-path"), loggerFormat)
		if err != nil {
			return err
		}

		// start web application and enter main loop

		return nil
	},
}
