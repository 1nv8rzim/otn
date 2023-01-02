package main

import (
	"github.com/1nv8rzim/otn/api"
	"github.com/1nv8rzim/otn/database"
)

func main() {
	defer database.Client.Close()

	api.Run()
}
