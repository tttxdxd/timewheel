package timewheel

import (
	"math/rand"
	"testing"
)

func BenchmarkDelayQueue(b *testing.B) {
	queue := newDelayQueue()
	for i:=0;i<100;i++{
		queue.Push(newTask(rand.Int63(),false,nil))
	}


	cases := []struct {
		name string
		N    int // the data size (i.e. number of existing timers)
	}{
		{"N-1m", 1000},
	}

	for _,c:=range cases{
		b.Run(c.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				queue.Push(newTask(rand.Int63(),false, nil))
				queue.Pop()
			}
		})
	}

	b.Log(queue)
}
