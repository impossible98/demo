package v1

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// import local packages
	"dragonfly/internal/service/user"
	"dragonfly/pkg/ecode"
	"dragonfly/pkg/format"
)

type User struct {
	Identity string `json:"identity" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

// Create User
func CreateUser(ctx *gin.Context) {
	request := User{
		Identity: uuid.NewString(),
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		format.HTTP(ctx, ecode.InvalidParams, err, nil)
		return
	}
	result, err := user.CreateUser(request.Identity, request.Username, request.Password, request.Nickname)
	if err != nil {
		format.HTTP(ctx, ecode.ErrorCreateUser, err, nil)
		return
	}
	format.HTTP(ctx, ecode.Success, nil, result)
}
