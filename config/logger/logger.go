package logger

import (
	"gin-cli/settings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(cfg *settings.LogConfig) error {
	//构造自定义的logger
	encoder := getEncoder()
	writeSyncer := getLogWriter(cfg.FileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	//Log Level：哪种级别的日志将被写入
	var level = new(zapcore.Level)
	err := level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return err
	}
	//使用配置项，构造一个zap core
	core := zapcore.NewCore(encoder, writeSyncer, level)
	//根据配置构造出我们自定义的logger
	//为了使调用者信息也包含在日志中，可以增加一个option：zap.AddCaller()
	log := zap.New(core, zap.AddCaller())
	// 替换zap库中全局的logger
	zap.ReplaceGlobals(log)
	return nil
}

func getEncoder() zapcore.Encoder {
	//如果对输出的日志的内容有要求，可以自定义配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //日期的输出格式
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	//Zap本身不支持切割归档日志文件,但是可以借助第三方Lumberjack
	//要在zap中加入Lumberjack支持，我们需要修改WriteSyncer代码
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  //日志文件的位置
		MaxSize:    maxSize,   //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: maxBackup, //保留旧文件的最大个数
		MaxAge:     maxAge,    //保留旧文件的最大天数
		Compress:   false,     //是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
