package api

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"dragonfly/global"
	"dragonfly/pkg/ecode"
	"dragonfly/pkg/format"
)

func Version(ctx *gin.Context) {
	format.HTTP(ctx, ecode.Success, nil, gin.H{
		"version": global.VERSION,
	})
}
