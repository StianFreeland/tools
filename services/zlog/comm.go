package zlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const moduleName = "zlog"

var logger *zap.Logger

func Start() {
	loadConfig()
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		zapcore.Lock(zapcore.AddSync(&writer{})),
		atomicLevel,
	)
	cores := []zapcore.Core{fileCore}
	if debug {
		stdoutCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(config),
			zapcore.Lock(wrap{os.Stdout}),
			atomicLevel,
		)
		cores = append(cores, stdoutCore)
	}
	logger = zap.New(
		zapcore.NewTee(cores...),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.PanicLevel),
	)
}

func Stop() {
	if err := logger.Sync(); err != nil {
		fmt.Println(moduleName, "sync, error:", err)
	}
}

func GetLevel() int64 {
	return int64(atomicLevel.Level())
}

func UpdateLevel(level int64) {
	atomicLevel.SetLevel(zapcore.Level(level))
}

func Fatal(prefix string, msg string, fields ...zap.Field) {
	logger.Fatal("[ "+prefix+" ] "+msg, fields...)
}

func Error(prefix string, msg string, fields ...zap.Field) {
	logger.Error("[ "+prefix+" ] "+msg, fields...)
}

func Warn(prefix string, msg string, fields ...zap.Field) {
	logger.Warn("[ "+prefix+" ] "+msg, fields...)
}

func Info(prefix string, msg string, fields ...zap.Field) {
	logger.Info("[ "+prefix+" ] "+msg, fields...)
}
