package datasource

import (
	"fmt"
	"gin-cli/settings"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var RedisDb *redis.Client

func InitRedis(cfg *settings.RedisConfig) error {
	rdb, err := NewRedis(cfg)
	RedisDb = rdb
	return err
}

// 返回Redis实例
func NewRedis(cfg *settings.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		zap.L().Error("redis连接出错", zap.Error(err))
		return nil, err
	}
	zap.L().Info("redis连接成功")
	return rdb, nil
}
