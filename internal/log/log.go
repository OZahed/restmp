package log

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LOG_LEVEL int8

const (
	L_DEBUG int8 = iota - 1
	L_INFO
	L_WARN
	L_ERR
)

var (
	Logger *zap.Logger
	Level  zap.AtomicLevel
)

const logLayout = "2006-01-02 15:04:05.000"

// init: Set NewProduction as default logger. Config depend on Logger instance
func Initialize(_ context.Context, l LOG_LEVEL) error {
	var err error
	Level = zap.NewAtomicLevelAt(zapcore.Level(l))
	Logger, err = zap.Config{
		Level:            Level,
		Development:      false,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			EncodeTime:     zapcore.TimeEncoderOfLayout(logLayout),
			EncodeDuration: zapcore.StringDurationEncoder,
			LevelKey:       "level",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			NameKey:        "name",
			EncodeName:     zapcore.FullNameEncoder,
			CallerKey:      "caller",
			EncodeCaller:   zapcore.FullCallerEncoder,
			MessageKey:     "msg",
			LineEnding:     zapcore.DefaultLineEnding,
		},
	}.Build()

	if err != nil {
		return err
	}
	return nil
}

func Finalize() error {
	return Logger.Sync()
}
