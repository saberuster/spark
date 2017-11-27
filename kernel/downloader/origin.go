package downloader

import (
	"net/http"
	"github.com/saberuster/spark/kernel/taskpool"
)

type OriginDownloader struct {
	client *http.Client
}

func (od *OriginDownloader) Do(t *taskpool.Task) (err error) {
	t.Response, err = od.client.Do(t.Request)
	return
}

func (od *OriginDownloader) Copy() Downloader {
	return &OriginDownloader{
		client: od.client,
	}
}

func NewDownloader(c *http.Client) Downloader {
	return &OriginDownloader{
		client: c,
	}
}
