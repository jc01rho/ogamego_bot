package Queue

import (
	"sync"
	"testing"
	"time"
)

func Test_test(t *testing.T) {
	q := Queue{ Items:make(chan Jobs, 100) ,C: sync.NewCond(new(sync.Mutex))}
	q.Set(Printtest,1)
	q.Set(Printtest,2)
	q.Set(Printtest,3)

	go q.DirectRun()
	go q.DirectRun()
	go q.DirectRun()

	go q.DirectRun()
	time.Sleep(time.Second*5)
	q.Set(Printtest,4)



}
