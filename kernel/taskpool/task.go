package taskpool

import "net/http"


type Task struct {
	Request  *http.Request
	Response *http.Response
}
