package main

import (
	"github.com/1nv8rzim/otu-paste/api"
	"github.com/1nv8rzim/otu-paste/database"
)

func main() {
	defer database.Client.Close()

	api.Run()
}
