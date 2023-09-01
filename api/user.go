package api

import (
	"github.com/gin-gonic/gin"
	"go-init/global"
	"go-init/models"
	"go-init/service"
)

var (
	userService = service.NewUserService()
)

func userToVo(user *models.User) *models.UserVo {
	return &models.UserVo{
		ID:          user.ID,
		UserName:    user.UserName,
		UserAccount: user.UserAccount,
		UserAvatar:  user.UserAvatar,
		Gender:      user.Gender,
		UserRole:    user.UserRole,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

// Login 用户登录
func Login(c *gin.Context) {
	var request models.UserLogin
	if err := c.ShouldBind(&request); err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}
	user, err := userService.Login(c, request.UserAccount, request.Password)
	if err != nil {
		global.ResponseFailed(c, global.ErrorInvalidParams)
		return
	}

	global.ResponseSuccess(c, userToVo(user))
}
