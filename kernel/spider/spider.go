package spider

import "github.com/saberuster/spark/kernel/taskpool"

type Spider struct {
	Task *taskpool.Task
	Result interface{}
}
