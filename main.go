package main

import (
	"gin-cli/settings"
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

	//3.初始化数据源

	//5.注册路由

	//6.启动服务
}
