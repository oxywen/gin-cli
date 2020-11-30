package initialize

import (
	"gin-server-cli/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// Redis 初始化redis
func Redis() *redis.Client {
	redisConfig := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.ZapLog.Error("redis connect ping failed, err:", zap.Any("err", err))
		return nil
	} else {
		global.ZapLog.Info("redis connect ping response:", zap.String("pong", pong))
		return client
	}
}
