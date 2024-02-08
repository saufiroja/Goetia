package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

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
