package global

import (
	"github.com/redis/go-redis/v9"
	"go-init/config"
	"gorm.io/gorm"
)

var (
	ServerConfig = &config.ServerConfig{}
	DB           = &gorm.DB{}
	CaChe        = &redis.Client{}
)
