package kernel

import (
	"github.com/saberuster/spark/kernel/downloader"
	"github.com/saberuster/spark/kernel/taskpool"
	"sync"
	"github.com/saberuster/spark/kernel/pipeline"
	"github.com/saberuster/spark/kernel/spider"
	"fmt"
	debug2 "github.com/saberuster/spark/common/debug"
	"os"
	"os/signal"
	"syscall"
)

var defaultConcurrent = 100

var (
	quit     chan bool
	relaunch bool
	debug    debug2.Debug
)

type Kernel struct {
	dlPool    downloader.PoolInterface //下载器
	taskPool  taskpool.Interface       //任务池
	pipelines []pipeline.Writer
}

func (k *Kernel) Run() {
	quit = make(chan bool)
	debug = true

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

//func lookPath() (argv0 string, err error) {
//	argv0, err = exec.LookPath(os.Args[0])
//	if nil != err {
//		return
//	}
//	if _, err = os.Stat(argv0); nil != err {
//		return
//	}
//	return
//}
