package OgameUtil

import "bitbucket.org/jc01rho/ogame"

func BuildTargetBuilding(bot *ogame.OGame, targetPlanet ogame.CelestialID, targetBuilding ogame.BaseOgameObj) {
	//bot.SendFleet()
	_ = bot.BuildBuilding(targetPlanet, targetBuilding.GetID())
}
