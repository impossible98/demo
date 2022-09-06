package user

import (
	// import local packages
	"dragonfly/internal/model"
)

// Create User
func CreateUser(identity string, username string, password string, nickname string) (*model.User, error) {
	user := model.User{
		Identity: identity,
		Username: username,
		Password: password,
		Nickname: nickname,
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
