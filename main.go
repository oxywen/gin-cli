package main

import (
	"gin-server-cli/core"
	"gin-server-cli/core/initialize"
	"gin-server-cli/global"
)

func main() {
	global.Viper = core.Viper("config.yaml")
	global.ZapLog = core.Zap()
	global.DbEngine = initialize.MySQL()
	global.RedisDb = initialize.Redis()
	core.RunHTTPServer()
}
