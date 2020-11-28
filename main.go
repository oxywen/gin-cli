package main

import (
	"gin-cli/logger"
	"gin-cli/settings"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
 * go web开发较通用的脚手架
 */
func main() {
	//1.加载配置
	err := settings.Init()
	if err != nil {
		//配置文件读取出错了就直接结束主程序
		return
	}

	//2.初始化日志
	logger.InitLogger(settings.Conf.LogConfig)

	//3.初始化数据源

	//5.注册路由
	r := gin.New()
	//gin框架中使用gin-zap中间件
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	//6.启动服务

}
