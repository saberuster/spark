package main

import (
	"github.com/saberuster/spark/kernel/taskpool"
	"net/http"
	"github.com/saberuster/spark/kernel"
	"github.com/saberuster/spark/kernel/downloader"
	"github.com/saberuster/spark/kernel/pipeline"
	"fmt"
	"github.com/saberuster/spark/kernel/spider"
)

func main() {
	go serverStart()
	dlPool := downloader.NewPool(downloader.DefaultDownLoader(), 100)
	taskPool := taskpool.NewTaskPool(100)

	req, err := http.NewRequest("Get", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	task := &taskpool.Task{Request: req}
	taskPool.Add(task)

	pipelines := make([]pipeline.Writer, 0, 1)
	pipelines = append(pipelines, &DebugPipeline{})

	k := kernel.New(dlPool, taskPool, pipelines)
	k.Run()
}

func serverStart() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", nil)

}

type DebugPipeline struct {
}

func (*DebugPipeline) Write(spider *spider.Spider) error {

	if r, ok := spider.Result.(*http.Response); ok {
		fmt.Println("收到数据:")
		buf := make([]byte, 2)
		r.Body.Read(buf)
		fmt.Printf("%s...", string(buf))
		r.Body.Close()
	}
	return nil
}
