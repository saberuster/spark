package taskpool

var pool = NewTaskPool(defaultPoolSize)

func Add(task *Task) bool {
	return pool.Add(task)
}

func Get() (t *Task, ok bool) {
	return pool.Get()
}

func Default() Interface {
	return pool
}
