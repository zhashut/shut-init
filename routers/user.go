package routers

import (
	"github.com/gin-gonic/gin"
	"go-init/api"
)

func InitUser(r *gin.RouterGroup) {
	u := r.Group("/user")
	{
		u.POST("/login", api.Login)
	}
}
