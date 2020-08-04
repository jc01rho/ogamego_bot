package Global

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/jc01rho/ogamego_bot/OGameBotRoutine"
	"github.com/jc01rho/ogamego_bot/Queue"
	"github.com/jc01rho/ogamego_bot/Scheduler"
	"gopkg.in/urfave/cli.v2"
	"strconv"
	"strings"
)

func InitEssentials(bot *ogame.OGame,c *cli.Context) {

	//OGameBotRoutine.OGameBotGlobal.Ogamebot = Test.GetBotForTest()
	//OGameBotRoutine.OGameBotGlobal.Class = OGameBotRoutine.OGameBotGlobal.Ogamebot.CharacterClass()

	//&cli.StringFlag{
	//	Name:    "main-coords",
	//	Usage:   "Main coords",
	//	Value:   "",
	//	EnvVars: []string{"MAIN_COORDS"},
	//},
	//	&cli.BoolFlag{
	//		Name:    "expedtiion-mode",
	//		Usage:   "Expedition mode",
	//		Value:   true,
	//		EnvVars: []string{"EXPEDITION_MODE"},
	//	},

	IsExpeditionMode = c.Bool("expedtiion-mode")
	stringCoords := strings.Split(c.String("main-coords"), ":")


	tempMainCoords := new(ogame.Coordinate)
	tempMainCoords.Galaxy, _ = strconv.ParseInt(stringCoords[0], 10, 64)
	tempMainCoords.System, _ =  strconv.ParseInt((stringCoords[1]),10,64)
	tempMainCoords.Position , _=  strconv.ParseInt((stringCoords[2]),10,64)
	typeTmp,_ :=  (strconv.ParseInt((stringCoords[3]),10,64))
	tempMainCoords.Type = ogame.CelestialType(typeTmp)
	MainCoords = tempMainCoords




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
