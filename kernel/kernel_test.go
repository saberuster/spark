package kernel

import (
	"testing"
	"net/http"
	"github.com/saberuster/spark/kernel/downloader"
	"github.com/saberuster/spark/kernel/taskpool"
	"github.com/saberuster/spark/kernel/spider"
	"fmt"
	"github.com/saberuster/spark/kernel/pipeline"
)

func TestKernel_Run(t *testing.T) {
	go serverStart()
	req, err := http.NewRequest("Get", "http://localhost:8080", nil)
	if err != nil {
		t.Fatal(err)
	}
	task := taskpool.Task{
		Request: req,
	}
	pool := taskpool.NewTaskPool(100)
	pool.Add(&task)

	fp := &firstPipeline{}
	sp := &secondPipeline{}

	k := &Kernel{
		downloader:    downloader.DefaultDownLoader(),
		pool:          pool,
		MaxConcurrent: 5,
		pipelines:     []pipeline.Writer{fp, sp},
	}
	k.Run()
}

func serverStart() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", nil)

}

type firstPipeline struct {
}

func (fp *firstPipeline) Write(spider *spider.Spider) error {
	bts := make([]byte, 2)
	spider.Task.Response.Body.Read(bts)
	spider.Result = string(bts)
	return nil
}

type secondPipeline struct {
}

func (sp *secondPipeline) Write(spider *spider.Spider) error {
	fmt.Printf("收到了数据'%s'", spider.Result)
	return nil
}
