package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func setupLogger(logLevel string, logDir string, formatter string) error {
	// set log output to local file if logPath specified, os.Stdout by default.
	if len(logDir) == 0 {
		fmt.Fprintln(os.Stdout, "use stdout as default log output")
		log.SetOutput(os.Stdout)
	} else {
		stat, err := os.Stat(logDir)
		if err != nil && os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "log-dir %s not exists", logDir)
			return err
		}

		if !stat.IsDir() {
			fmt.Fprintf(os.Stderr, "log-dir %s is not a valid directory", logDir)
			return errors.New("log-dir is not a directory")
		}

		logPath := filepath.Join(logDir, "web.log")
		logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}
		log.SetOutput(logFile)
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

	return nil
}
