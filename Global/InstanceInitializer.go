package Global

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/jc01rho/ogamego_bot/OGameBotRoutine"
	"github.com/jc01rho/ogamego_bot/Queue"
	"github.com/jc01rho/ogamego_bot/Scheduler"
)

func InitEssentials(bot *ogame.OGame) {

	//OGameBotRoutine.OGameBotGlobal.Ogamebot = Test.GetBotForTest()
	//OGameBotRoutine.OGameBotGlobal.Class = OGameBotRoutine.OGameBotGlobal.Ogamebot.CharacterClass()
	OGameBotRoutine.OGameBotGlobal.Ogamebot = bot
	OGameBotRoutine.OGameBotGlobal.Class = OGameBotRoutine.OGameBotGlobal.Ogamebot.CharacterClass()
	OGameBotRoutine.OGameBotGlobal.DetermineMainPlanet()
	OGameBotRoutine.OGameBotGlobal.BuildRessSkipList = hashset.New()

	Queue.InitQueue()
	//Bot.InitBot()

	//Queue.InitQueue()
	Scheduler.InitScheduler()
	Scheduler.DefaultJobs()

}
