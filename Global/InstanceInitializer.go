package Global

import (
	"github.com/jc01rho/ogamego_bot/OGameBot"
	"github.com/jc01rho/ogamego_bot/Queue"
	"github.com/jc01rho/ogamego_bot/Scheduler"
	"github.com/jc01rho/ogamego_bot/Test"
)

func InitEssentials() {

	OGameBot.OGameBotGlobal.Ogamebot = Test.GetBotForTest()

	Queue.InitQueue()
	//Bot.InitBot()

	Queue.InitQueue()
	Scheduler.InitScheduler()
	Scheduler.DefaultJobs()

}
