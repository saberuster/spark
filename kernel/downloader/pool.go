package downloader

import (
	"sync"
	"fmt"
)

type PoolInterface interface {
	Get() (Downloader, error)
	Free(Downloader)
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

func NewPool(dl Downloader, cc int) *Pool {
	dls := make([]Downloader, cc)
	c := make(chan bool, cc)
	for n := 0; n < cc; n++ {
		dls = append(dls, dl.Copy())
		c <- true
	}

	return &Pool{
		downloaders: dls,
		count:       c,
		mu:          sync.Mutex{},
	}
}
