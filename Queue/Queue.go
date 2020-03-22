package Queue

import (
	"fmt"
	"github.com/emirpasic/gods/trees/binaryheap"
	log "github.com/sirupsen/logrus"

	"path/filepath"
	"reflect"
	"runtime"
	"sync"
)

//TODO : change to priority queue
const (
	CriticalPriority = iota + 0

	MajorPriority
	DefaultPriority
	MinorPriority
)

func JobsPrioritySorter(a, b interface{}) int {
	// Type assertion, program will panic if this is not respected
	c1 := a.(Jobs)
	c2 := b.(Jobs)

	switch {
	case c1.Priority > c2.Priority:
		return 1
	case c1.Priority < c2.Priority:
		return -1
	default:
		return 0
	}
}

type Jobs struct {
	Funcs    interface{}   // Map for the function task store
	Fparams  []interface{} // Map for function and  params of function
	Priority int
}
type Queue struct {
	Heap *binaryheap.Heap

	Counter   int
	Counsume  *sync.Cond
	Produce   *sync.Cond
	SizeLimit int
}

var JobQueue *Queue = nil

// 값을 저장
func (q *Queue) Set(priority int, value interface{}, params ...interface{}) {
	//defer q.Counsume.Signal() // will wake up a popper
	//q.Counsume.L.Lock()

	//defer q.Counsume.L.Unlock()

	fucntionpointer := runtime.FuncForPC(reflect.ValueOf(value).Pointer())
	f, l := fucntionpointer.FileLine(fucntionpointer.Entry())
	log.Debug("Trying Queue add [", filepath.Base(f), ":", l, "] ", fucntionpointer.Name())
	job := Jobs{
		Priority: priority,
		Funcs:    value,
		Fparams:  params,
	}
	//q.Items <- job

	//
	//job = <- q.Items
	q.Produce.L.Lock()
	if q.Heap.Size() >= q.SizeLimit {
		q.Produce.Wait()
	}
	q.Counter++
	q.Heap.Push(job)
	q.Counsume.Signal()
	q.Produce.L.Unlock()

	log.Debug("Add Complete. Current Queue Size is ", q.Counter)

}

func (q *Queue) DirectRun() {
	//q.Counsume.L.Lock()
	//defer q.Counsume.L.Unlock()
	//
	//for len(q.Items) == 0 {
	//	q.Counsume.Wait()
	//}

	q.Counsume.L.Lock()
	if q.Heap.Size() == 0 {
		q.Counsume.Wait()
	}
	var tmpitem, result = q.Heap.Pop()
	q.Counter--

	if q.Heap.Size() < q.SizeLimit {
		q.Produce.Signal()
	}

	q.Counsume.L.Unlock()

	item := tmpitem.(Jobs)

	if result == false {
		return
	}

	f := reflect.ValueOf(item.Funcs)
	if len(item.Fparams) != f.Type().NumIn() && !reflect.TypeOf(item.Funcs).IsVariadic() {
		//return nil, errors.New("the number of params is not matched")
	}
	in := make([]reflect.Value, len(item.Fparams))
	for k, param := range item.Fparams {
		in[k] = reflect.ValueOf(param)
	}

	fucntionpointer := runtime.FuncForPC(f.Pointer())

	fc, l := fucntionpointer.FileLine(fucntionpointer.Entry())

	log.Info("Queue Task Calling [", filepath.Base(fc), ":", l, "] ", runtime.FuncForPC(f.Pointer()).Name())

	f.Call(in)

}

func Printtest(a int) {
	fmt.Print(a)
}

func InitQueue() {
	newQ := new(Queue)
	newQ.Counsume = sync.NewCond(new(sync.Mutex))
	newQ.Produce = sync.NewCond(new(sync.Mutex))
	newQ.SizeLimit = 100
	//newQ.Items = make(chan Jobs, newQ.SizeLimit)
	newQ.Counter = 0

	newQ.Heap = binaryheap.NewWith(JobsPrioritySorter)

	JobQueue = newQ
}
