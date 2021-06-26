package main

import (
	"bootstrap/booty"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	log.SetFormatter(Formatter)

	var log = logrus.New()

	log.Info("Bootstraping with booty and away we go!!!")
	booty.Execute()
}
