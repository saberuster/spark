package downloader

import "testing"

func TestNewPool(t *testing.T) {
	p := NewPool(DefaultDownLoader(), 1)

	d, err := p.Get()
	if err != nil {
		t.Fatal(err)
	}

	p.Free(d)
}
