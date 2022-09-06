package main

import (
	// import local packages
	"dragonfly/internal/database"
	"dragonfly/internal/model"
	"dragonfly/internal/server"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		panic(err)
	}
	model.DB = db
	server.InitServer()
}
