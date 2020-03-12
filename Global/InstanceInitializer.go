package Global

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/jc01rho/ogamego_bot/OGameBot"
	"github.com/jc01rho/ogamego_bot/Queue"
	"github.com/jc01rho/ogamego_bot/Scheduler"
)

func InitEssentials(bot *ogame.OGame) {

	//OGameBot.OGameBotGlobal.Ogamebot = Test.GetBotForTest()
	//OGameBot.OGameBotGlobal.Class = OGameBot.OGameBotGlobal.Ogamebot.CharacterClass()
	OGameBot.OGameBotGlobal.Ogamebot = bot
	OGameBot.OGameBotGlobal.Class = OGameBot.OGameBotGlobal.Ogamebot.CharacterClass()
	OGameBot.OGameBotGlobal.DetermineMainPlanet()
	OGameBot.OGameBotGlobal.BuildRessSkipList = hashset.New()

	Queue.InitQueue()
	//Bot.InitBot()

	//Queue.InitQueue()
	Scheduler.InitScheduler()
	Scheduler.DefaultJobs()

}
