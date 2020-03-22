package OGameBotRoutine

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/emirpasic/gods/sets"
)

var OGameBotGlobal OGameBot

type OGameBot struct {
	Ogamebot              *ogame.OGame
	MainPlanetCoord       ogame.Coordinate
	MainPlanetCelestitial ogame.CelestialID
	IsMainPlanetMoon      bool

	Class ogame.CharacterClass

	BuildRessSkipList sets.Set

	Universe string
	Username string
	Password string
	Language string
}
