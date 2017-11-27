package downloader

import (
	"github.com/saberuster/spark/kernel/taskpool"
)

type Downloader interface {
	Do(task *taskpool.Task) error
	Copy() Downloader
}
