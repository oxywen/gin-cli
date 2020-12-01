package main

import (
	"gin-server-cli/core"
	"gin-server-cli/core/application"
	"gin-server-cli/core/initialize"
)

func main() {
	application.Viper = core.Viper("config.yaml")
	application.Log = core.Zap()
	application.DbEngine = initialize.Gorm()
	application.Redis = initialize.Redis()
	core.RunHTTPServer()
}
