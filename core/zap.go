package core

import (
	"fmt"
	"gin-server-cli/global"
	"gin-server-cli/utils"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.Config.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.Config.Zap.Director)
		_ = os.Mkdir(global.Config.Zap.Director, os.ModePerm)
	}
	var level zapcore.Level
	// 初始化配置文件的Level
	switch global.Config.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(level), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore(level))
	}
	if global.Config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.Config.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.Config.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.Config.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.Config.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.Config.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.Config.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(level zapcore.Level) zapcore.Core {
	var writer zapcore.WriteSyncer
	var err error
	switch global.Config.Zap.RotateType {
	case "lumberjack":
		writer = lumberJackWriteSyncer()
	case "file-rotatelogs":
		writer, err = getRotateLogsWrite()
	default:
		writer, err = getRotateLogsWrite()
	}
	if err != nil {
		fmt.Printf("get write syncer failed err:%v", err.Error())
		return nil
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

func lumberJackWriteSyncer() zapcore.WriteSyncer {
	var lumberJackWriter = &lumberjack.Logger{
		Filename:   global.Config.Zap.Director + "/" + global.Config.Zap.FileName, //日志文件的路径
		MaxSize:    global.Config.Zap.MaxSize,                                     //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: global.Config.Zap.MaxBackups,                                  //保留旧文件的最大个数
		MaxAge:     global.Config.Zap.MaxAge,                                      //保留旧文件的最大天数
		Compress:   false,                                                         //是否压缩/归档旧文件
	}
	if global.Config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackWriter))
	}
	return zapcore.AddSync(lumberJackWriter)
}

func getRotateLogsWrite() (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(global.Config.Zap.Director, "%Y-%m-%d.log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if global.Config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.Config.Zap.Prefix + "2006-01-02 15:04:05.000"))
}
