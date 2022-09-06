package router

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"dragonfly/internal/server/api"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/api/version", api.Version)
	return router
}
