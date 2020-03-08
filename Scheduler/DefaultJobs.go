package Scheduler

import (
	"github.com/jc01rho/ogamego_bot/OGameBot"
	"github.com/jc01rho/ogamego_bot/Queue"
	"time"
)

func DefaultJobs() {

	//Scheduler.Every(2).Seconds().Do(Queue.JobQueue.DirectRun)

	//Scheduler.Every(1).Second().Do(JustPrint)

	Queue.JobQueue.Set(func() { OGameBot.OGameBotGlobal.MaintainLCCountStep(-1) })
	Scheduler.Every(1).Hour().Do(Queue.JobQueue.Set, OGameBot.OGameBotGlobal.BuildNextRess)
	Scheduler.Every(1).Hour().Do(Queue.JobQueue.Set, OGameBot.OGameBotGlobal.BuildNextRess)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			Queue.JobQueue.DirectRun()
		}
	}()

}
