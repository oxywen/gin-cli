package application

import (
	"gin-server-cli/core/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//本来是global包的，但是嫌弃global不好看，换成application了
//存放全局的对象，例如gorm engine，redis client等
var (
	DbEngine *gorm.DB
	Redis    *redis.Client
	Config   config.Application
	Viper    *viper.Viper
	Log      *zap.Logger
)
