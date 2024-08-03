package logger

import (
	"time"

	"github.com/RSheremeta/web-crawler/config"
	"github.com/sirupsen/logrus"
)

func NewDefaultLogger() *logrus.Entry {
	return logrus.NewEntry(logrus.StandardLogger())
}

func NewLogger(cfg *config.Config) *logrus.Entry {
	logger := NewDefaultLogger()
	if cfg == nil || cfg.Logger == nil {
		return logger
	}

	logger.Logger.SetLevel(logrus.InfoLevel)

	if cfg.Logger.DebugLevel {
		logger.Logger.SetLevel(logrus.DebugLevel)
	}

	logger.Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.DateTime,
	})

	return logger
}
