package database

import (
	// import third-party packages
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// import local packages
	"dragonfly/internal/model"
	"dragonfly/pkg/path"
)

func InitDatabase() (*gorm.DB, error) {
	if !path.PathExists("data") {
		os.Mkdir("data", os.ModePerm)
	}
	database, err := gorm.Open(sqlite.Open("data/dragonfly.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = database.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}
	return database, nil
}
