package Scheduler

import (
	"github.com/jc01rho/gocron"
	"math/rand"
	"time"
)

var Scheduler *gocron.Scheduler

func InitScheduler() {
	Scheduler = gocron.NewScheduler()
	go func() {
		<-Scheduler.Start()
	}()

}

func SleepRandomSeconds(waitTime int) {
	var randmin int64 = rand.Int63n(15)
	time.Sleep(time.Second * time.Duration(randmin))
}

func SleepRandomMinutes(waitTime int) {
	var randmin int64 = rand.Int63n(15)
	time.Sleep(time.Minute * time.Duration(randmin))
}
