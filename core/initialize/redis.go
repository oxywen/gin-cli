package initialize

import (
	"gin-server-cli/core/application"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// Redis 初始化redis
func Redis() *redis.Client {
	redisConfig := application.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		application.Log.Error("redis connect ping failed, err:", zap.Any("err", err))
		return nil
	} else {
		application.Log.Info("redis connect ping response:", zap.String("pong", pong))
		return client
	}
}
