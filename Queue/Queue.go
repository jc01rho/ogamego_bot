package Queue

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"path/filepath"
	"reflect"
	"runtime"
	"sync"
)

//TODO : change to priority queue

type Jobs struct {
	Funcs   interface{}   // Map for the function task store
	Fparams []interface{} // Map for function and  params of function
}
type Queue struct {
	Items   chan Jobs
	Counter int
	C       *sync.Cond
}

var JobQueue *Queue = nil

// 값을 저장
func (q *Queue) Set(value interface{}, params ...interface{}) {
	//defer q.C.Signal() // will wake up a popper
	//q.C.L.Lock()
	//defer q.C.L.Unlock()

	fucntionpointer := runtime.FuncForPC(reflect.ValueOf(value).Pointer())
	f, l := fucntionpointer.FileLine(fucntionpointer.Entry())
	log.Debug("Trying Queue add [", filepath.Base(f), ":", l, "] ", fucntionpointer.Name())
	job := Jobs{

		Funcs:   value,
		Fparams: params,
	}
	q.Items <- job
	q.Counter++
	log.Trace("Add Complete. Current Queue Size is ", q.Counter)

}

// 값을 꺼내기
func (q *Queue) Get() Jobs {
	log.Info("Get")
	q.Counter--
	return <-q.Items
}

func (q *Queue) DirectRun() {
	//q.C.L.Lock()
	//defer q.C.L.Unlock()
	//
	//for len(q.Items) == 0 {
	//	q.C.Wait()
	//}

	var item Jobs
	q.Counter--
	item = <-q.Items

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
	newQ.C = sync.NewCond(new(sync.Mutex))
	newQ.Items = make(chan Jobs, 100)
	newQ.Counter = 0

	JobQueue = newQ
}

func test() {
	q := Queue{Items: make(chan Jobs, 100), C: sync.NewCond(new(sync.Mutex))}
	q.Set(Printtest)
	q.Set(Printtest)
	q.Set(Printtest)

	a := q.Get()
	_ = a
	fmt.Print("1")

}
