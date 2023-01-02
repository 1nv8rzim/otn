package database

import (
	"context"
	"log"
	"time"

	"github.com/1nv8rzim/otu-paste/config"
	"github.com/1nv8rzim/otu-paste/ent"
	"github.com/1nv8rzim/otu-paste/ent/endpoint"
	"github.com/1nv8rzim/otu-paste/helpers"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Client *ent.Client
	Ctx    context.Context = context.Background()
)

func init() {
	client, err := ent.Open("sqlite3", "file:database.sqlite?_loc=auto&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	Client = client

	// Run the auto migration tool.
	if err := Client.Schema.Create(Ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func RemoveOldEndpoints() {
	oldEndpoints, err := Client.Endpoint.Query().Where(endpoint.CreatedAtLTE(time.Now().Add(-config.EXPIRATION))).All(Ctx)
	if err != nil {
		return
	}
	for _, endpoint := range oldEndpoints {
		Client.Endpoint.DeleteOne(endpoint).Exec(Ctx)
	}
}

func CreateEndpoint(content string) (string, error) {
	endpoint, err := Client.Endpoint.Create().
		SetContent(content).
		SetURI(helpers.GenerateURI()).
		Save(Ctx)

	return endpoint.URI, err
}
