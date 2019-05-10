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

var logDirFlag = cli.StringFlag{
	Name:   "log-dir",
	EnvVar: "LOG_DIR",
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

var usdtRpcAddrFlag = cli.StringFlag{
	Name:   "usdt-rpc-addr",
	Value:  "localhost:18332",
	EnvVar: "USDT_RPCADDR",
	Usage:  "[NOTICE] testnet and mainnet have different default port",
}

var usdtRpcUserFlag = cli.StringFlag{
	Name:   "usdt-rpc-user",
	Value:  "foo",
	EnvVar: "USDT_RPCUSER",
}

var usdtRpcPassFlag = cli.StringFlag{
	Name:   "usdt-rpc-pass",
	Value:  "usdtpass",
	EnvVar: "USDT_PRCPASS",
	Usage:  "password can be generate through scripts/rcpauth.py",
}

var usdtPropertyIdFlag = cli.IntFlag{
	Name:   "usdt-property-id",
	Value:  2,
	EnvVar: "USDT_PROPERTY_ID",
	Usage:  "property id of usdt, default is 2",
}

var ethRpcAddrFlag = cli.StringFlag{
	Name:   "eth-rpc-addr",
	Value:  "http://192.168.0.101:8545",
	EnvVar: "ETH_RPCADDR",
}

var ethWalletDirFlag = cli.StringFlag{
	Name:   "eth-wallet-dir",
	Value:  "/data/eth-wallet",
	EnvVar: "ETH_WALLET_DIR",
}

var ethAccountPasswordFlag = cli.StringFlag{
	Name:   "eth-account-password",
	Value:  "password",
	EnvVar: "ETH_ACCOUNT_PASSWORD",
}

var ethAccountFlag = cli.StringFlag{
	Name:   "eth-account-path",
	Value:  "/data/eth-accounts.json",
	EnvVar: "ETH_ACCOUNT_PATH",
}

var envFlag = cli.StringFlag{
	Name:   "env",
	Value:  "production",
	EnvVar: "ENV",
}

var backendsFlag = cli.StringFlag{
	Name:   "backends",
	Value:  "btc,usdt,eth",
	EnvVar: "BACKENDS",
}
