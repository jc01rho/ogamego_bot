package OGameBot

import "bitbucket.org/jc01rho/ogame"

var OGameBotGlobal OGameBot

type OGameBot struct {
	Ogamebot *ogame.OGame

	Class ogame.CharacterClass

	Universe string
	Username string
	Password string
	Language string
}
