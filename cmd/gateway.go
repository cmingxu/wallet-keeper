package main

import (
	"github.com/cmingxu/wallet-keeper/api"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var gateCmd = cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Flags: []cli.Flag{
		httpAddrFlag,
		btcRpcAddrFlag,
		btcRpcUserFlag,
		btcRpcPassFlag,
	},
	Usage: "serve api gateway",
	Action: func(c *cli.Context) error {
		log.Infof("starting api gateway with addr: %s", c.String("http-listen-addr"))

		apiServer, err := api.NewApiServer(c.String("http-listen-addr"))
		if err != nil {
			return nil
		}

		apiServer.InitBtcClient(
			c.String("btc-rpc-addr"), // host
			c.String("btc-rpc-user"), // user
			c.String("btc-rpc-pass"), // password
		)

		// start accepting http requests
		return apiServer.HttpListen()
	},
}
