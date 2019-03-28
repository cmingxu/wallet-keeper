package main

import (
	"log"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Commands = []cli.Command{
		WebCmd,
	}
	app.Flags = []cli.Flag{
		logLevelFlag,
		logPathFlag,
	}

	app.Before(func() *cli.Context {
	})

}
