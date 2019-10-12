package application

import (
	"github.com/go-elastic-search/src/api/controller"
)

func mapEndpointsToControllers() {
	Router.GET("/ping", controller.Ping)
}
