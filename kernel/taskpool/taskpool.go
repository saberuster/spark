package taskpool

import "sync"

var defaultPoolSize = 10000

type Interface interface {
	Get() (*Task, bool)
	Add(*Task) bool
}

type TaskPool struct {
	tasks      []*Task
	count      int
	MaxTaskNum int
	mu         sync.Mutex
}

func (tp *TaskPool) Get() (t *Task, ok bool) {
	tp.mu.Lock()
	defer tp.mu.Unlock()
	if tp.count == 0 {
		return nil, false
	}
	t = tp.tasks[len(tp.tasks)-1]
	ok = true
	tp.tasks = tp.tasks[0:len(tp.tasks)-1]
	tp.count--
	return
}

func (tp *TaskPool) Add(task *Task) bool {
	tp.mu.Lock()
	defer tp.mu.Unlock()
	if tp.count >= tp.MaxTaskNum {
		return false
	}
	tp.tasks = append(tp.tasks, task)
	tp.count++
	return true
}

func NewTaskPool(poolSize int) Interface {
	if poolSize == 0 {
		poolSize = defaultPoolSize
	}
	return &TaskPool{
		tasks:      make([]*Task, 0, poolSize),
		MaxTaskNum: poolSize,
		count:      0,
		mu:         sync.Mutex{},
	}
}
