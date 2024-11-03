package utils

import (
	    "github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func initLogger() {
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetLevel(logrus.InfoLevel)
}