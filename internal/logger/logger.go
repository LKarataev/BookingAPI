package logger

import (
	"fmt"
	"log"
)

type Logger interface {
	InitLogger()
	Info(format string, v ...any)
	Errorf(format string, v ...any)
}

type apiLogger struct {
	stdLogger *log.Logger
}

func NewApiLogger() *apiLogger {
	var logger apiLogger
	logger.InitLogger()
	return &logger
}

func (l *apiLogger) InitLogger() {
	l.stdLogger = log.Default()
}

func (l *apiLogger) Errorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.stdLogger.Printf("[Error]: %s\n", msg)
}

func (l *apiLogger) Info(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.stdLogger.Printf("[Info]: %s\n", msg)
}
