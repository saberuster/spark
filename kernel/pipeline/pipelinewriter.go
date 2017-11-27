package pipeline

import (
	"github.com/saberuster/spark/kernel/spider"
)

type Writer interface {
	Write(spider *spider.Spider) error
}
