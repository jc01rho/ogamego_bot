package Scheduler

import (
	"github.com/jc01rho/ogamego_bot/OGameBot"
	"github.com/jc01rho/ogamego_bot/Queue"
	"time"
)

func DefaultJobs() {

	//Scheduler.Every(2).Seconds().Do(Queue.JobQueue.DirectRun)

	//Scheduler.Every(1).Second().Do(JustPrint)
	Scheduler.Every(2).Seconds().Do(Queue.JobQueue.Set, OGameBot.OGameBotGlobal.BuildNextRess)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			Queue.JobQueue.DirectRun()
		}
	}()

}
