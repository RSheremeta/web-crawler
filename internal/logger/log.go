package logger

import "github.com/sirupsen/logrus"

func NewLoggerInstance() *logrus.Entry {
	logger := logrus.NewEntry(logrus.StandardLogger())
	logger.Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	return logger
}
