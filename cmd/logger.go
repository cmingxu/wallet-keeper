package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func setupLogger(logLevel string, logPath string, formatter string) error {
	// set log output to local file if logPath specified, os.Stdout by default.
	if len(logPath) != 0 {
		logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}
		log.SetOutput(logFile)
	} else {
		log.SetOutput(os.Stdout)
	}

	// set log as json or text format
	if formatter == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	}

	// report caller method name in log message
	log.SetReportCaller(true)
	return nil
}
