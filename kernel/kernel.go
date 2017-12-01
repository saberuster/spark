package kernel

import (
	"sync"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/saberuster/spark/kernel/downloader"
	"github.com/saberuster/spark/kernel/taskpool"
	"github.com/saberuster/spark/kernel/pipeline"
	"github.com/saberuster/spark/kernel/spider"
)

var defaultConcurrent = 100

var (
	quit     chan struct{} //使用struct{}可以最小化占用的内存
	relaunch bool
)

type Kernel struct {
	dlPool    downloader.PoolInterface //下载器
	taskPool  taskpool.Interface       //任务池
	pipelines []pipeline.Writer
}

func (k *Kernel) Run() {
	quit = make(chan struct{})

	go signalHandler()

	wg := sync.WaitGroup{}
	for {
		t, ok := k.taskPool.Get()
		if !ok {
			break
		}
		dl, err := k.dlPool.Get()
		if err != nil {
			fmt.Println(err)
			panic(nil)
		}
		wg.Add(1)
		go func() {
			err := dl.Do(t)
			if err != nil {
				fmt.Println(err)
				return
			}

			sp := &spider.Spider{Task: t, Result: t.Response}
			for _, pip := range k.pipelines {
				err = pip.Write(sp)
				if err != nil {
					fmt.Println(err)
					break
				}
			}
			k.dlPool.Free(dl)
			wg.Done()
		}()
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

func signalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1)
	for sig := range c {
		if sig == syscall.SIGUSR1 {
			relaunch = true
		}

		close(quit)
		break
	}
}

