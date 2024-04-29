package logger

import (
	"fmt"
	"log"
	"sync"
)

type LoggerInterface interface {
	Init(logger *log.Logger)
	Info(format string, v ...any)
	Errorf(format string, v ...any)
}

type logger struct {
	stdLogger *log.Logger
}

func (l *logger) Init(logger *log.Logger) {
	l.stdLogger = logger
}

func (l *logger) Info(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.stdLogger.Printf("[Info]: %s\n", msg)
}

func (l *logger) Errorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.stdLogger.Printf("[Error]: %s\n", msg)
}

var loggerImpl LoggerInterface
var lock sync.Mutex

func GetLogger() LoggerInterface {
	lock.Lock()
	defer lock.Unlock()

	if loggerImpl == nil {
		loggerImpl = new(logger)
		loggerImpl.Init(log.Default())
	}
	return loggerImpl
}
