package Scheduler

import (
	"github.com/jc01rho/ogamego_bot/OGameBot"
	"github.com/jc01rho/ogamego_bot/Queue"
	"time"
)

func DefaultJobs() {

	//Scheduler.Every(2).Seconds().Do(Queue.JobQueue.DirectRun)

	//Scheduler.Every(1).Second().Do(JustPrint)

	Scheduler.Every(6).Hours().Do(Queue.JobQueue.Set, OGameBot.OGameBotGlobal.BuildNextRess)
	//Queue.JobQueue.Set(func() { OGameBot.OGameBotGlobal.MaintainLCCountStep(-1) })
	//Queue.JobQueue.Set(func() { OGameBot.OGameBotGlobal.CollectRessRoutine() })
	//Queue.JobQueue.Set(func() { OGameBot.OGameBotGlobal.BuildNextRess() })
	Scheduler.Every(12).Hour().Do(Queue.JobQueue.Set, OGameBot.OGameBotGlobal.MaintainLCCountStep, -1)
	Scheduler.Every(24).Hour().Do(Queue.JobQueue.Set, OGameBot.OGameBotGlobal.CollectRessRoutine)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			Queue.JobQueue.DirectRun()
		}
	}()

}
