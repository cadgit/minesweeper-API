package main

import (
	"github.com/sirupsen/logrus"
	"minesweeper-API/api"
)

func main() {
	log := logrus.StandardLogger()
	log.Infoln("Starting API ...")
	err := api.Start(log)

	if err != nil {
		log.WithError(err).Fatalln("Start API fail")
	}
}