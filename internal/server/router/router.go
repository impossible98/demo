package router

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"dragonfly/internal/server/api"
	"dragonfly/internal/server/api/v1"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/api/version", api.Version)
	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/user/create", v1.CreateUser)
	}
	return router
}
