package datasource

import "gin-cli/settings"

//初始化所有的数据源，例如mysql redis
func InitDatasource(cfg *settings.DatasourceConfig) error {
	err := InitMySQL(cfg.MySQLConfig)
	err = InitRedis(cfg.RedisConfig)
	return err
}

//配置自动迁移
func AutoMigrateConfig() {

}
