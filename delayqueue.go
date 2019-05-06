package timewheel

import (
	"math"
	"sync/atomic"
	"time"
)

type task struct {
	job      func()
	id       int
	interval int64
	delay    int64
	repeat   bool
	times    int64
	_start   int64
}

var autoId = int32(0)

func newTask(interval int64, repeat bool, job func()) *task {
	now := time.Now().UnixNano()
	return &task{
		id:       int(atomic.AddInt32(&autoId, 1)),
		interval: interval,
		delay:    interval,
		repeat:   repeat,
		job:      job,
		_start:   now,
	}
}

func (t *task) reset() {
	t.delay = t.interval*(t.times+1) + t._start
}

type delayQueue []*task

func newDelayQueue() *delayQueue {
	return &delayQueue{newTask(math.MinInt64, true, nil)}
}

func (q delayQueue) Len() int {
	return len(q)
}

func (q delayQueue) Less(i, j int) bool {
	return q[i].delay < q[j].delay
}

func (q delayQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *delayQueue) Push(v *task) {
	*q = append(*q, v)
	q.shiftUp(q.Len() - 1)
}

func (q *delayQueue) Pop() *task {
	item := q.Peek()

	if q.Len() > 1 {
		q.shiftDown()
		*q = (*q)[0 : q.Len()-1]
	}

	return item
}

func (q *delayQueue) Peek() *task {
	if q.Len() > 1 {
		return (*q)[1]
	} else {
		return nil
	}
}

// shiftUp 上浮
func (q *delayQueue) shiftUp(i int) {
	item := (*q)[i]
	for ; item.delay < (*q)[i/2].delay; i = i / 2 {
		(*q)[i] = (*q)[i/2]
	}
	(*q)[i] = item
}

// shiftDown 下浮
func (q *delayQueue) shiftDown() {
	l := q.Len() - 1
	i, child := 1, 2
	for child < l {
		if child+1 < l && q.Less(child+1, child) {
			child++
		}
		if q.Less(child, l) {
			(*q)[i] = (*q)[child]
		} else {
			break
		}

		i, child = child, child*2
	}
	(*q)[i] = (*q)[l]
}
