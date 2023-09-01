package service

import (
	"context"
	"go-init/models"
	"go-init/respository/db"
)

type UserService struct {
	db.BaseDao
}

func NewUserService() *UserService {
	return &UserService{}
}

// Login 用户登录
func (s *UserService) Login(ctx context.Context, account, password string) (*models.User, error) {
	return nil, nil
}
