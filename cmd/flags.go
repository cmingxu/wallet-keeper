package main

import (
	"github.com/urfave/cli"
)

var httpAddrFlag = cli.StringFlag{
	Name:   "http-listen-addr",
	Value:  "0.0.0.0:8000",
	Usage:  "http address of web application",
	EnvVar: "HTTP_LISTEN_ADDR",
}

var logLevelFlag = cli.StringFlag{
	Name:   "log-level",
	Value:  "info",
	Usage:  "default log level",
	EnvVar: "LOG_LEVEL",
}

var logPathFlag = cli.StringFlag{
	Name:   "log-path",
	EnvVar: "LOG_PATH",
}

var btcAddr = cli.StringFlag{
	Name:   "btc-addr",
	Value:  "http://localhost:8585",
	EnvVar: "BTC_ADDR",
}
