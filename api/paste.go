package api

import (
	"github.com/1nv8rzim/otn/database"
	"github.com/1nv8rzim/otn/ent/endpoint"
	"github.com/gin-gonic/gin"
)

func paste(ctx *gin.Context) {
	id := ctx.Param("id")

	database.RemoveOldEndpoints()

	endpoint, err := database.Client.Endpoint.Query().Where(endpoint.URI(id)).Only(database.Ctx)
	if err != nil {
		ctx.String(404, "note not found")
		return
	}

	database.Client.Endpoint.DeleteOne(endpoint).Exec(database.Ctx)

	ctx.String(200, endpoint.Content)
}
