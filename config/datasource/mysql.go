package datasource

import (
	"fmt"
	"gin-cli/settings"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DbEngine *gorm.DB

func InitMySQL(cfg *settings.MySQLConfig) error {
	db, err := NewDbEngine(cfg)
	DbEngine = db
	return err
}

func NewDbEngine(cfg *settings.MySQLConfig) (*gorm.DB, error) {
	//拿到数据库的配置信息拼接成数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	//创建数据库引擎
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//如果数据库连接失败
		zap.L().Error("数据库连接失败", zap.String("dsn", dsn), zap.Error(err))
		return nil, err
	}
	//输出日志
	zap.L().Info("数据库连接成功", zap.String("dsn", dsn))
	//设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("数据库连接池获取失败", zap.Error(err))
		return nil, err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(cfg.MaxLifetime))
	zap.L().Info("数据库连接池设置完毕")
	return db, nil
}
