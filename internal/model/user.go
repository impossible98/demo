package model

import (
	// import third-party packages
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Identity string `gorm:"column:identity;not null" json:"identity"` // Favorites表的id
	Username string `gorm:"column:username;not null;unique" json:"username"`
	Password string `gorm:"column:password;not null" json:"password"`
	Nickname string `gorm:"column:nickname;not null" json:"nickname"`
}

func (User) TableName() string {
	return "user"
}
