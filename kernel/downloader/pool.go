package downloader

import (
	"sync"
	"net/http"
	"fmt"
)

type PoolInterface interface {
	Get() (d Downloader, err error)
	Free(d Downloader)
}

type Pool struct {
	downloaders []Downloader
	count       chan bool
	mu          sync.Mutex
}

func (p *Pool) Get() (d Downloader, err error) {
	<-p.count
	p.mu.Lock()
	defer p.mu.Unlock()
	endPos := len(p.downloaders) - 1
	fmt.Println(endPos)
	d = p.downloaders[endPos]
	p.downloaders = p.downloaders[0:endPos]
	return
}

func (p *Pool) Free(d Downloader) {
	p.mu.Lock()
	p.downloaders = append(p.downloaders, d)
	p.mu.Unlock()
	p.count <- true
	return
}

func NewPool() *Pool {
	client := &http.Client{}
	dls := make([]Downloader, 0, 10)
	c := make(chan bool, 10)
	for n := 0; n < 10; n++ {
		dls = append(dls, NewDownloader(client))
		c <- true
	}

	return &Pool{
		downloaders: dls,
		count:       c,
		mu:          sync.Mutex{},
	}
}
