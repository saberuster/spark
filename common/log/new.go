package log

import (
	"io"
	"log"
)

func New(out io.Writer, prefix string, flag int, level Level) *Logger {
	return &Logger{
		log:   log.New(out, prefix, flag),
		level: level,
	}
}

func NewFromLog(logger *log.Logger, level Level) *Logger {
	return &Logger{
		log:   logger,
		level: level,
	}
}
