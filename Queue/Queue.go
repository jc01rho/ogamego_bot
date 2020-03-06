package Queue

import (
	"fmt"
	"github.com/jc01rho/ogamego_bot/Logger"
	"reflect"
	"sync"
)

// 정수현 chan으로 Queue정의
type Jobs struct {
	Funcs   interface{}   // Map for the function task store
	Fparams []interface{} // Map for function and  params of function
}
type Queue struct {
	Items chan Jobs
	C     *sync.Cond
}

var JobQueue *Queue = nil

// 값을 저장
func (q *Queue) Set(value interface{}, params ...interface{}) {
	//defer q.C.Signal() // will wake up a popper
	//q.C.L.Lock()
	//defer q.C.L.Unlock()
	Logger.Logger.Info("Set")
	job := Jobs{

		Funcs:   value,
		Fparams: params,
	}
	q.Items <- job

}

// 값을 꺼내기
func (q *Queue) Get() Jobs {
	Logger.Logger.Info("Get")
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
	item = <-q.Items

	f := reflect.ValueOf(item.Funcs)
	if len(item.Fparams) != f.Type().NumIn() && !reflect.TypeOf(item.Funcs).IsVariadic() {
		//return nil, errors.New("the number of params is not matched")
	}
	in := make([]reflect.Value, len(item.Fparams))
	for k, param := range item.Fparams {
		in[k] = reflect.ValueOf(param)
	}

	f.Call(in)

}

func Printtest(a int) {
	fmt.Print(a)
}

func InitQueue() {
	newQ := new(Queue)
	newQ.C = sync.NewCond(new(sync.Mutex))
	newQ.Items = make(chan Jobs, 100)

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
