package core

import (
	"flag"
	"fmt"
	"gin-server-cli/core/application"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	//命令行传参会覆盖掉默认的config.yaml
	var config string
	//调用者指定配置文件路径
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" {
		//命令行没有传参，那就看开发者是否指定默认的
		if len(path) == 0 {
			config = "config.yaml"
			fmt.Printf("用户未指定配置文件路径，将使用默认的配置文件路径%v\n", config)
		} else {
			config = path[0]
			fmt.Printf("用户未指定配置文件路径，将使用开发者指定的配置文件路径%v\n", config)
		}
	} else {
		//命令行传入了配置文件路径，那就是用命令行传入的
		fmt.Printf("正在使用用户指定的配置文件的路径%v\n", config)
	}
	//构建viper实例
	v := viper.New()
	//传入配置文件路径
	v.SetConfigFile(config)
	//读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("viper has an error reading the configuration file:%s\n", err))
	}
	//开始监听配置文件的修改
	v.WatchConfig()
	//当配置文件被修改了之后触发回调
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("the configuration file has been modified:%s\n", e.Name)
		//同时更新全局的配置文件
		if err := v.Unmarshal(&application.Config); err != nil {
			fmt.Println(err)
		}
	})
	//把配置文件反序列化为结构体
	if err := v.Unmarshal(&application.Config); err != nil {
		fmt.Printf("viper has an error unmarshal:%s\n", err.Error())
	}
	return v
}
