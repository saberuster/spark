package spark

import (
	"syscall"
	"os"
	"os/exec"
	"github.com/saberuster/spark/kernel/downloader"
	"github.com/saberuster/spark/kernel/taskpool"
	"github.com/saberuster/spark/kernel/spider"
	"fmt"
	"github.com/saberuster/spark/kernel/pipeline"
	"github.com/saberuster/spark/kernel"
	"net/http"
	debug2 "github.com/saberuster/spark/common/debug"
)

var (
	quit     chan bool
	relaunch bool
	debug    debug2.Debug
)

func Run() {
	quit = make(chan bool)
	debug = true
	debug.Println("开始运行")
	go signalHandler() //监听信号

	run()

	if relaunch {
		//重新启动服务
		argv0, err := lookPath()
		if nil != err {
			return
		}

		err = syscall.Exec(argv0, os.Args, os.Environ())
		if err != nil {
			return
		}
	}
}

func lookPath() (argv0 string, err error) {
	argv0, err = exec.LookPath(os.Args[0])
	if nil != err {
		return
	}
	if _, err = os.Stat(argv0); nil != err {
		return
	}
	return
}

func run() {

	debug.Println("初始化下载器")
	dlPool := downloader.NewPool()

	debug.Println("初始化任务池")
	taskPool := taskpool.NewTaskPool(100)

	debug.Println("添加task")
	req, err := http.NewRequest("Get", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	task := &taskpool.Task{Request: req}
	taskPool.Add(task)

	debug.Println("初始化pipeline")
	pipelines := make([]pipeline.Writer, 0, 1)
	pipelines = append(pipelines, &DebugPipeline{})

	k := kernel.New(dlPool, taskPool, pipelines)
	k.Run()
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
