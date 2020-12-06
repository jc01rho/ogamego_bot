package OgameUtil

import "bitbucket.org/jc01rho/ogame"

func BuildTargetID(bot *ogame.OGame, targetPlanet ogame.CelestialID, targetBuilding ogame.BaseOgameObj, nbr int64) {
	//bot.SendFleet()

	_ = bot.Build(targetPlanet, targetBuilding.GetID(),nbr)
}
