package zlog

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var debug bool
var atomicLevel zap.AtomicLevel

func loadConfig() {
	debug = viper.GetBool("zap.debug")
	atomicLevel = zap.NewAtomicLevelAt(zapcore.Level(viper.GetInt64("zap.level")))
}
