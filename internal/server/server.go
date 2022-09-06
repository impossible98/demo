package server

import (
	// import local packages
	"dragonfly/internal/server/router"
)

func InitServer() {
	server := router.Router()
	server.Run()
}
