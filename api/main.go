package api

import (
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func init() {
	router := gin.Default()
	Router = router

	routes()
}

func routes() {
	Router.POST("/", index_post)
	Router.GET("/", index_get)
	Router.GET("/:id", paste)

}

func Run() {
	Router.Run(":80")
}
