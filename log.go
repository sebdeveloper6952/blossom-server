package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLog(level string) (*zap.Logger, error) {
	logLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	if level == "DEBUG" {
		logLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	cfg := zap.Config{
		Encoding:         "json",
		Level:            logLevel,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, err := cfg.Build()

	return logger, err
}
