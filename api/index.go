package api

import (
	"github.com/1nv8rzim/otu-paste/database"
	"github.com/1nv8rzim/otu-paste/helpers"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	content := ctx.PostForm("content")

	if content == "" {
		ctx.String(400, "empty request")
		return
	}

	database.RemoveOldEndpoints()
	uri, err := database.CreateEndpoint(content)
	if err != nil {
		ctx.String(500, "error creating endpoint")
	}

	ctx.String(200, helpers.GenerateURL(uri))
}
