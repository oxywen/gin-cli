package global

import (
	"gin-server-cli/core/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//存放全局的对象，例如gorm engine，redis client等
var (
	DbEngine *gorm.DB
	RedisDb  *redis.Client
	Config   config.Application
	Viper    *viper.Viper
	ZapLog   *zap.Logger
)
