package timewheel

import "time"

type timeWheel struct {
	register   chan *task
	unregister chan int
	wheel      *delayQueue
	flagQueue  map[int]struct{}
}

func NewTimeWheel() *timeWheel {
	return &timeWheel{
		register:   make(chan *task),
		unregister: make(chan int),
		wheel:      newDelayQueue(),
		flagQueue:  make(map[int]struct{}),
	}
}

func (t *timeWheel) addTask(timeout int64, repeat bool, job func()) int {
	task := newTask(timeout*int64(time.Millisecond), repeat, job)
	t.register <- task
	return task.id
}

func (t *timeWheel) act(task *task) {
	if _, ok := t.flagQueue[task.id]; ok {
		t.wheel.Pop()
		delete(t.flagQueue, task.id)
		return
	}

	t.wheel.Pop()
	if task.repeat {
		task.reset()
		t.wheel.Push(task)
	}

	task.job()
	task.times++
}

func (t *timeWheel) run() {
	for {
		task := t.wheel.Peek()
		if task == nil {
			select {
			case task := <-t.register:
				t.wheel.Push(task)
			case id := <-t.unregister:
				t.flagQueue[id] = struct{}{}
			}
		} else {
			now := time.Now().UnixNano()
			timeout := task.delay - now
			if timeout <= 0 {
				t.act(task)
			} else {
				select {
				case task := <-t.register:
					t.wheel.Push(task)
				case id := <-t.unregister:
					t.flagQueue[id] = struct{}{}
				case <-time.After(time.Duration(timeout)):
					t.act(task)
				}
			}
		}
	}
}

func (t *timeWheel) Start() {
	go t.run()
}

func (t *timeWheel) Stop(id int) {
	t.flagQueue[id] = struct{}{}
}

func (t *timeWheel) SetTimeout(timeout int64, job func()) int {
	return t.addTask(timeout, false, job)
}

func (t *timeWheel) SetInterval(timeout int64, job func()) int {
	return t.addTask(timeout, true, job)
}
