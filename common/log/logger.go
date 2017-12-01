package log

import "fmt"

func (l *Logger) Info(format string, args ...interface{}) {
	if l.level&LVINFO == 1 {
		l.log.Println(fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	if l.level&LVDEBUG == 1 {
		l.log.Println(fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Warning(format string, args ...interface{}) {
	if l.level&LVWARN == 1 {
		l.log.Println(fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Error(format string, args ...interface{}) {
	if l.level&LVERROR == 1 {
		l.log.Panicln(fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	if l.level&LVFATAL == 1 {
		l.log.Fatal(fmt.Sprintf(format, args...))
	}
}
