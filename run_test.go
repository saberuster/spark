package spark

import (
	"testing"
	"net/http"
)

func TestRun(t *testing.T) {
	go serverStart()
	Run()
}

func serverStart() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", nil)

}
