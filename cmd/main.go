package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Universal BTC/USDT wallet gateway"
	app.Version = Version
	app.Commands = []cli.Command{
		gateCmd,
	}

	app.Flags = []cli.Flag{
		logLevelFlag,
		logDirFlag,
		envFlag,
	}

	app.Before = func(c *cli.Context) error {
		return setupLogger(c.String("log-level"), c.String("log-dir"), "json")
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
