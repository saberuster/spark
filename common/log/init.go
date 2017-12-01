package log

import (
	"log"
	"os"
)

var defaultLogger = &Logger{
	log:   log.New(os.Stderr, "", log.LstdFlags),
	level: LVDEFAULT,
}

func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args)
}

func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args)
}

func Warning(format string, args ...interface{}) {
	defaultLogger.Warning(format, args)
}

func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args)
}

func Fatal(format string, args ...interface{}) {
	defaultLogger.Fatal(format, args)
}
