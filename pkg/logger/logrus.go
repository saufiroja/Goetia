//go:generate mockgen -destination ../../mocks/mock_logger.go -package mocks github.com/saufiroja/cqrs/pkg/logger ILogger
package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

type ILogger interface {
	StartLogger(fileName string, nameFunc string) *logrus.Entry
}

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)
	log.SetOutput(logger.Writer())
	logger.SetOutput(io.MultiWriter(os.Stdout))

	return &Logger{logger}
}

func (l *Logger) StartLogger(fileName string, nameFunc string) *logrus.Entry {
	return l.WithFields(logrus.Fields{
		"file": fileName,
		"func": nameFunc,
	})
}
