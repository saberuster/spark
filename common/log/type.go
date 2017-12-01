package log

import "log"

type Level = int

const (
	LVFATAL   = 1 << 0
	LVERROR   = LVFATAL & 1 << 1
	LVWARN    = LVERROR & 1 << 2
	LVINFO    = LVWARN & 1 << 3
	LVDEBUG   = LVINFO & 1 << 4
	LVDEFAULT = LVERROR
)

type Logger struct {
	log   *log.Logger
	level Level
}


