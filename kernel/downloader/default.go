package downloader

import (
	"github.com/saberuster/spark/kernel/taskpool"
)

func Do(task *taskpool.Task) (err error) {
	task.Response, err = defaultClient.Do(task.Request)
	return
}

func DefaultDownLoader() Downloader {
	return NewDownloader(defaultClient)
}
