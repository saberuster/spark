package downloader

import (
	"sync"
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
	d = p.downloaders[endPos]
	p.downloaders = p.downloaders[0:endPos]
	return
}

func (p *Pool) Free(d Downloader) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.count) == cap(p.count) {
		return
	}
	p.count <- true
	p.downloaders = append(p.downloaders, d)
	return
}

func NewPool(dl Downloader, cc int) *Pool {
	dls := make([]Downloader, 0, cc)
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
