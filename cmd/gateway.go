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
		usdtRpcAddrFlag,
		usdtRpcUserFlag,
		usdtRpcPassFlag,
	},
	Usage: "serve api gateway",
	Action: func(c *cli.Context) error {

		apiServer, err := api.NewApiServer(c.String("http-listen-addr"))
		if err != nil {
			return nil
		}

		log.Infof("connecting to btc rpc addr: %s", c.String("btc-rpc-addr"))
		err = apiServer.InitBtcClient(
			c.String("btc-rpc-addr"), // host
			c.String("btc-rpc-user"), // user
			c.String("btc-rpc-pass"), // password
		)
		if err != nil {
			log.Error(err)
			return err
		}

		// Check btc/usdt rpc call connectivity
		err = apiServer.KeeperCheck()
		if err != nil {
			log.Error(err)
			return err
		}

		log.Infof("starting api gateway with addr: %s", c.String("http-listen-addr"))
		// start accepting http requests
		return apiServer.HttpListen()
	},
}
