package Scheduler

import (
	"github.com/jc01rho/ogamego_bot/OGameBotRoutine"
	"github.com/jc01rho/ogamego_bot/Queue"
	log "github.com/sirupsen/logrus"
	"time"
)

func HeaertBeat() {
	log.Info("HeartBeat Bump!")
}

func DefaultJobs() {
	//	//Scheduler.Every(2).Hours().From(gocron.NextTick()).Do(Queue.JobQueue.Set, Queue.CriticalPriority, HeaertBeat)
	//Scheduler.Every(10).Seconds().From(gocron.NextTick()).Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.BuildNextRess)

	//Scheduler.Every(1).Day().At("08:01").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.BuildNextRess)

	Scheduler.Every(1).Day().At("08:00").Do(Queue.JobQueue.Set, Queue.CriticalPriority, OGameBotRoutine.RestartLogic)

	Scheduler.Every(20).Minutes().At("08:04").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.BuildDefs)
	Scheduler.Every(25).Minutes().At("08:03").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.MaintainLCCountStep, int64(-1))
	Scheduler.Every(11).Hours().At("08:06").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.CollectRessRoutine)
	Scheduler.Every(2).Hours().At("11:01").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.BuildNextRess)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			Queue.JobQueue.DirectRun()
		}
	}()

}
