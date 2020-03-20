package Scheduler

import (
	"github.com/jc01rho/ogamego_bot/OGameBot"
	"github.com/jc01rho/ogamego_bot/Queue"
	log "github.com/sirupsen/logrus"
	"time"
)

func HeaertBeat() {
	log.Info("HeartBeat Bump!")
}

func DefaultJobs() {

	//Scheduler.Every(2).Seconds().Do(Queue.JobQueue.DirectRun)

	//Scheduler.Every(1).Second().Do(JustPrint)

	Scheduler.Every(10).Seconds().Do(Queue.JobQueue.Set, HeaertBeat)
	Scheduler.Every(6).Hours().Do(Queue.JobQueue.Set, OGameBot.OGameBotGlobal.BuildNextRess)
	//Queue.JobQueue.Set(func() { OGameBot.OGameBotGlobal.MaintainLCCountStep(-1) })
	//Queue.JobQueue.Set(func() { OGameBot.OGameBotGlobal.CollectRessRoutine() })
	//Queue.JobQueue.Set(func() { OGameBot.OGameBotGlobal.BuildNextRess() })
	Scheduler.Every(12).Hours().Do(Queue.JobQueue.Set, OGameBot.OGameBotGlobal.MaintainLCCountStep, -1)
	Scheduler.Every(24).Hours().Do(Queue.JobQueue.Set, OGameBot.OGameBotGlobal.CollectRessRoutine)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			Queue.JobQueue.DirectRun()
		}
	}()

}
