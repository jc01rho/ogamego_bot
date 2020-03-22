package Queue

import (
	"sync"
	"testing"
	"time"
)

func Test_test(t *testing.T) {
	q := Queue{Items: make(chan Jobs, 100), Counsume: sync.NewCond(new(sync.Mutex))}
	q.Set(4, Printtest, 1)
	q.Set(3, Printtest, 2)
	q.Set(2, Printtest, 3)

	go q.DirectRun()
	go q.DirectRun()
	go q.DirectRun()

	go q.DirectRun()
	time.Sleep(time.Second * 5)
	q.Set(1, Printtest, 4)

}
