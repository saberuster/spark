package kernel

import (
	"github.com/saberuster/spark/kernel/downloader"
	"github.com/saberuster/spark/kernel/taskpool"
	"sync"
	"github.com/saberuster/spark/kernel/pipeline"
	"github.com/saberuster/spark/kernel/spider"
	"fmt"
)

var defaultConcurrent = 100

type Kernel struct {
	dlPool    downloader.PoolInterface //下载器
	taskPool  taskpool.Interface       //任务池
	pipelines []pipeline.Writer
}

func (k *Kernel) Run() {
	wg := sync.WaitGroup{}
	for {
		t, ok := k.taskPool.Get()
		if !ok {
			break
		}
		fmt.Println("获取到任务")
		dl, err := k.dlPool.Get()
		fmt.Println("获取到下载器")
		if err != nil {
			fmt.Println(err)
			panic(nil)
		}
		wg.Add(1)
		go func(dl downloader.Downloader) {
			defer func() {
				k.dlPool.Free(dl)
				wg.Done()
			}()
			err := dl.Do(t)
			if err != nil {
				fmt.Println(err)
				return
			}
			sp := &spider.Spider{Task: t, Result: t.Response}
			for _, pip := range k.pipelines {
				pip.Write(sp)
			}
		}(dl)
	}

	wg.Wait()
}

func New(dlPool downloader.PoolInterface, taskPool taskpool.Interface, pipelines []pipeline.Writer) *Kernel {
	return &Kernel{
		dlPool:    dlPool,
		taskPool:  taskPool,
		pipelines: pipelines,
	}
}
