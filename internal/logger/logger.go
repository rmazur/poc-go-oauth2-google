package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

/*
InitializeZapCustomLogger Funtion initializes a logger using uber-go/zap package in the application.
*/
func InitializeZapCustomLogger() {

	config := zap.NewProductionConfig()
	config.Encoding = "json"
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	config.OutputPaths = []string{viper.GetString("logger-output-path"), "stdout"}
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.CallerKey = "file"
	config.EncoderConfig.MessageKey = "msg"
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	Log, _ = config.Build()
	defer Log.Sync()
}
