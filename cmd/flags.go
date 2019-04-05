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

var btcRpcAddrFlag = cli.StringFlag{
	Name:   "btc-rpc-addr",
	Value:  "192.168.0.101:8332",
	EnvVar: "BTC_RPCADDR",
	Usage:  "[NOTICE] testnet and mainnet have different default port",
}

var btcRpcUserFlag = cli.StringFlag{
	Name:   "btc-rpc-user",
	Value:  "foo",
	EnvVar: "BTC_RPCUSER",
}

var btcRpcPassFlag = cli.StringFlag{
	Name:   "btc-rpc-pass",
	Value:  "qDDZdeQ5vw9XXFeVnXT4PZ--tGN2xNjjR4nrtyszZx0=",
	EnvVar: "BTC_PRCPASS",
	Usage:  "password can be generate through scripts/rcpauth.py",
}

var envFlag = cli.StringFlag{
	Name:   "env",
	Value:  "production",
	EnvVar: "ENV",
}
