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



	//Scheduler.Every(90).Minutes().From(gocron.NextTick()).Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.MaintainLCCountStep, int64(-1))

	Scheduler.Every(1).Day().At("08:00").Do(Queue.JobQueue.Set, Queue.CriticalPriority, OGameBotRoutine.RestartLogic)
	Scheduler.Every(1).Day().From(gocron.NextTick()).Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.BuildNextRess)

	//Scheduler.Every(1).Day().At("08:01").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.BuildNextRess)
	Scheduler.Every(1).Day().At("12:01").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.BuildNextRess)
	Scheduler.Every(1).Day().At("16:01").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.BuildNextRess)
	Scheduler.Every(1).Day().At("18:01").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.BuildNextRess)
	Scheduler.Every(1).Day().At("20:01").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.BuildNextRess)
	Scheduler.Every(1).Day().At("22:01").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.BuildNextRess)
	

	//Scheduler.Every(2).Hours().From(gocron.NextTick()).Do(Queue.JobQueue.Set, Queue.CriticalPriority, HeaertBeat)



	Scheduler.Every(1).Day().At("08:05").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.CollectRessRoutine)
	Scheduler.Every(90).Minutes().At("08:05").Do(Queue.JobQueue.Set, Queue.DefaultPriority, OGameBotRoutine.OGameBotGlobal.MaintainLCCountStep, int64(-1))

	go func() {
		for {
			time.Sleep(5 * time.Second)
			Queue.JobQueue.DirectRun()
		}
	}()

}
