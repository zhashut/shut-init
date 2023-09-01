package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-init/global"
	"go.uber.org/zap"
)

// InitConfig 初始化配置
func InitConfig() {
	configPrefix := "config"
	configFileName := fmt.Sprintf("./%s.yaml", configPrefix)
	v := viper.New()
	// 文件路径设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息：%v", global.ServerConfig)
	// viper的功能-动态监控变化
	_ = v.WriteConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件发生变化: %s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&global.ServerConfig)
		zap.S().Infof("配置信息：%v", global.ServerConfig)
	})
	zap.S().Infof("%v", global.ServerConfig)
}
