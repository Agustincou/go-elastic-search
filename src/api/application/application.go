package application

import (
	"github.com/gin-gonic/gin"
)

// Router : router endpoints
var (
	Router *gin.Engine
)

// Start : run the app
func Start() {
	configureRouter()
	mapEndpointsToControllers()
	Router.Run(":8080")
}

func configureRouter() {
	Router = gin.Default()
}