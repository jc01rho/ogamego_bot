package OGameBot

import "bitbucket.org/jc01rho/ogame"

var OGameBotGlobal OGameBot

type OGameBot struct {
	Ogamebot *ogame.OGame

	Universe string
	Username string
	Password string
	Language string
}
