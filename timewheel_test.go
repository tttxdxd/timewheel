package timewheel

import (
	"log"
	"testing"
	"time"
)

func TestDelayQueue(t *testing.T) {

	//queue := newDelayQueue()
	//data := []int64{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//for i := 0; i < 10; i++ {
	//	queue.Push(newTask(data[i],false, nil))
	//}
	//result := []int64{}
	//for queue.Peek() != nil {
	//	task :=queue.Pop()
	//	if task==nil{
	//		t.Error("nil")
	//	}
	//	result = append(result, task.delay)
	//}
	//
	//for i, n := range result {
	//	if i > 0 && result[i-1] > n {
	//		t.Error("error")
	//		break
	//	}
	//}
	//t.Log(result)

	timewheel:=NewTimeWheel()
	timewheel.Start()
	i:=0
	timewheel.SetTimeout(1000, func() {
		log.Println("tete")
	})
	now:=time.Now()
	timewheel.SetInterval(3000, func() {
		i++
		log.Println("interval-1 : ",i,"  ",time.Now().Sub(now).String())
	})
	j:=0
	timewheel.SetInterval(3000, func() {
		j++
		log.Println("interval-2 : ",j,"  ",time.Now().Sub(now).String())
	})

	time.Sleep(time.Minute*20)

}
