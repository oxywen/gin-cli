package core

import (
	"fmt"
	"gin-server-cli/core/application"
	"gin-server-cli/core/initialize"
)

func RunHTTPServer() {
	router := initialize.Routers()
	address := fmt.Sprintf(":%d", application.Config.System.Port)
	router.Run(address)
	//server := http.Server{
	//	Addr:           address,
	//	Handler:        router,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//application.ZapLog.Info("server run success on ", zap.String("address", address))
	//application.ZapLog.Error(server.ListenAndServe().Error())
}
