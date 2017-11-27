package downloader

import (
	"testing"
	"net/http"
	"fmt"
	"github.com/saberuster/spark/kernel/queue"
	"github.com/saberuster/spark/kernel/taskpool"
)

func TestDo(t *testing.T) {
	go serverStart()
	req, err := http.NewRequest("Get", "http://localhost:8080", nil)
	if err != nil {
		t.Fatal(err)
	}
	task := &taskpool.Task{Request: req}

	err = Do(task)
	if err != nil {
		t.Fatal(err)
	}
	defer task.Response.Body.Close()
	r := make([]byte, 2)
	task.Response.Body.Read(r)

	if string(r) != "ok" {
		t.Fatal(fmt.Sprintf("获取到的数据为 %s", string(r)))
	}

}

func serverStart() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", nil)

}
