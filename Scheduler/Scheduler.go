package Scheduler

import (
	"github.com/jc01rho/gocron"
)

var Scheduler *gocron.Scheduler

func InitScheduler() {
	Scheduler = gocron.NewScheduler()
	go func() {
		<-Scheduler.Start()
	}()

}
