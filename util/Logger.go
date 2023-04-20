package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func ConfigZap(logFile []string) {
	var (
		cfg zap.Config
		err error
	)

	cfg = zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding:    "json",
		OutputPaths: logFile,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:     "timestamp",
			EncodeTime:  zapcore.RFC3339NanoTimeEncoder,
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
	}

	Logger, err = cfg.Build()
	if err != nil {
		os.Exit(3)
	}

	return
}

func LogInfo(data []zap.Field) {
	Logger.Info("", data...)
}

func LogError(data []zap.Field) {
	Logger.Error("", data...)
}
