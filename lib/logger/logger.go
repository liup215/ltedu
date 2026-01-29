package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger, _ = zap.NewProduction()
)

func InitLogger(logCfg *Config) {
	var err error
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	lv := zap.InfoLevel
	switch strings.ToLower(logCfg.Level) {
	case "debug":
		lv = zap.DebugLevel
	case "info":
		lv = zap.InfoLevel
	case "warn", "warning":
		lv = zap.WarnLevel
	case "error":
		lv = zap.ErrorLevel
	default:
		lv = zap.InfoLevel
	}

	if logCfg.Filename == "" {
		cfg.Encoding = "console"
		if logCfg.LogEncoding != "console" {
			cfg.Encoding = "json"
		}
		cfg.Level.SetLevel(lv)

		paths := []string{"stdout"}
		cfg.ErrorOutputPaths = paths
		cfg.OutputPaths = paths
		Logger, err = cfg.Build()
		if err != nil {
			Logger.Fatal("zap build", zap.Error(err))
		}
		return
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logCfg.Filename,
		MaxSize:    logCfg.MaxSizeMB, // megabytes
		MaxBackups: logCfg.MaxBackups,
		MaxAge:     logCfg.MaxDays, // days
		Compress:   logCfg.Compress,
	})

	enc := zapcore.NewConsoleEncoder(cfg.EncoderConfig)
	if logCfg.LogEncoding != "console" {
		enc = zapcore.NewJSONEncoder(cfg.EncoderConfig)
	}
	core := zapcore.NewCore(
		enc,
		w,
		lv,
	)

	Logger = zap.New(
		core,
		zap.AddCaller(),
		// zap.AddCallerSkip(1),
	)

}
