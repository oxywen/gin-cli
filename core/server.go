package core

import (
	"fmt"
	"gin-server-cli/core/initialize"
	"gin-server-cli/global"
)

func RunHTTPServer() {
	router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.Config.System.Port)
	router.Run(address)
	//server := http.Server{
	//	Addr:           address,
	//	Handler:        router,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//global.ZapLog.Info("server run success on ", zap.String("address", address))
	//global.ZapLog.Error(server.ListenAndServe().Error())
}
