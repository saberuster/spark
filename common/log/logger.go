package log

import (
	"log"
	"io"
)

type Logger struct {
	log *log.Logger
}

func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{
		log: log.New(out, prefix, flag),
	}
}

func (l *Logger) Error(err error) {
	l.log.Println(err)
}


func (l *Logger) Info() {

}

func (l *Logger) Warning() {

}






