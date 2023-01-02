package main

import (
	"time"

	"github.com/1nv8rzim/otn/api"
	"github.com/1nv8rzim/otn/config"
	"github.com/1nv8rzim/otn/database"
)

var Ticker *time.Ticker = time.NewTicker(config.CLEANINTERVAL)

func main() {
	defer database.Client.Close()
	go clean()
	api.Run()
}

func clean() {
	for {
		<-Ticker.C
		database.RemoveOldEndpoints()
	}
}
