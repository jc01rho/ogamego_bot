package Scheduler

import (
	"github.com/jc01rho/gocron"
	"github.com/jc01rho/ogamego_bot/OGameBotRoutine"
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

	//Queue.JobQueue.Set(func() { OGameBotRoutine.OGameBotGlobal.MaintainLCCountStep(-1) })
	//Queue.JobQueue.Set(func() { OGameBotRoutine.OGameBotGlobal.CollectRessRoutine() })
	//Queue.JobQueue.Set(func() { OGameBotRoutine.OGameBotGlobal.BuildNextRess() })

	//Scheduler.Every(1).Second().Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.Ogamebot.IsUnderAttack)

	Scheduler.Every(1).Day().At("08:00").Do(Queue.JobQueue.Set, Queue.CriticalPriority, OGameBotRoutine.RestartLogic)
	Scheduler.Every(3).Hours().From(gocron.NextTick()).Do(Queue.JobQueue.Set, Queue.CriticalPriority, HeaertBeat)
	Scheduler.Every(6).Hours().From(gocron.NextTick()).Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.BuildNextRess)

	Scheduler.Every(30).Minutes().From(gocron.NextTick()).Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.MaintainLCCountStep, int64(-1))

	Scheduler.Every(24).Hours().From(gocron.NextTick()).Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.CollectRessRoutine)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			Queue.JobQueue.DirectRun()
		}
	}()

}
