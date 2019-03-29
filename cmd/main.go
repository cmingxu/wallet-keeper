package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		webCmd,
		jobCmd,
	}

	app.Flags = []cli.Flag{
		logLevelFlag,
		logPathFlag,
		envFlag,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
