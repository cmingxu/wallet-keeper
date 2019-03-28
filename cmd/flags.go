package main

import (
	"github.com/urfave/cli"
)

var httpAddrFlag cli.StringFlag = cli.StringFlag{
	Name:   "http-listen-addr",
	Value:  "0.0.0.0:8000",
	Usage:  "http address of web application",
	EnvVar: "HTTP_LISTEN_ADDR",
}

var logLevelFlag cli.StringFlag = cli.StringFlag{
	Name:   "log-level",
	Value:  "info",
	Usage:  "default log level",
	EnvVar: "LOG_LEVEL",
}

var logPathFlag cli.StringFlag = cli.StringFlag{
	Name:   "log-path",
	Value:  "/var/log/jex/web.log",
	EnvVar: "LOG_PATH",
}
