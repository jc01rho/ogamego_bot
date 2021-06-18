package OgameUtil

import "github.com/0xE232FE/ogame.mod"

func BuildTargetID(bot *ogame.OGame, targetPlanet ogame.CelestialID, targetBuilding ogame.BaseOgameObj, nbr int64) {
	//bot.SendFleet()

	_ = bot.Build(targetPlanet, targetBuilding.GetID(),nbr)
}
