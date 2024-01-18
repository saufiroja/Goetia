package logger

import "github.com/sirupsen/logrus"

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)
	return logger
}
