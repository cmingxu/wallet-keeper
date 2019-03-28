package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "web"
	app.Flags = []cli.Flag{
		httpAddrFlag,
		logLevelFlag,
		logPathFlag,
	}

	app.Usage = "jex backend web application"
	app.Action = func(c *cli.Context) error {
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Hello World")
}
