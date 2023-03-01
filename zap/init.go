package zap

import (
	"fmt"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZap() *zap.Logger {
	core := zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.DebugLevel)
	return zap.New(core, zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 路径[异步输出]
func getLogWriter() *zapcore.BufferedWriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("./storage/logs/log.%s.log", time.Now().Format("2006-01-02")),
		MaxSize:    10,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 7,     // 保留旧文件的最大个数
		MaxAge:     30,    // 保留旧文件的最大天数
		Compress:   false, // 是否压缩/归档旧文件
	}
	return &zapcore.BufferedWriteSyncer{
		WS:   zapcore.AddSync(lumberJackLogger),
		Size: 4096,
	}
}
