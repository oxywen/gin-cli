package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//全部变量，储存读取到的所有配置信息
var Conf = new(Config)

//定义结构体和配置文件相对应
type Config struct {
	*AppConfig        `mapstructure:"app"`
	*ServerConfig     `mapstructure:"server"`
	*LogConfig        `mapstructure:"log"`
	*DatasourceConfig `mapstructure:"datasource"`
}

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	MachineID int64  `mapstructure:"machine_id"`
	StartTime string `mapstructure:"start_time"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type DatasourceConfig struct {
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxLifetime  int    `mapstructure:"max_life_time"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}

/**
 * 使用viper读取配置文件并进行序列化
 */
func Init() error {
	// 指定项目根目录下的配置文件
	viper.SetConfigFile("./config.yaml")
	// 查找并读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper has an error reading the configuration file，err -> %v\n", err)
		return err
	}
	//把配置文件反序列化为结构体
	err = viper.Unmarshal(Conf)
	if err != nil {
		fmt.Printf("viper has an error unmarshal，err -> %v\n", err)
		return err
	}
	viper.WatchConfig() //开始监听配置文件的修改
	//当配置文件被修改了之后触发回调
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("the configuration file has been modified...")
		//同时更新全局Conf
		err := viper.Unmarshal(Conf)
		if err != nil {
			fmt.Printf("viper has an error unmarshal，err -> %v\n", err)
		}
	})
	return nil
}
