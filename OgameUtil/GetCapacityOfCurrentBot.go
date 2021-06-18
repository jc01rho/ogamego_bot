package OgameUtil

import "github.com/0xE232FE/ogame.mod"

func GetCapacityOfCurrentBot(bot *ogame.OGame, baseship ogame.BaseShip) int64 {

	//ogame.LargeCargo.GetCargoCapacity(currentResearch, bot.Ogamebot.GetServer().Settings.EspionageProbeRaids != 0 , bot.Class == ogame.Collector)
	return baseship.GetCargoCapacity(bot.GetCachedResearch(), bot.GetServer().Settings.EspionageProbeRaids != 0, bot.CharacterClass() == ogame.Collector, false)

}

func GetCapacityOfCurrentBotOfPlanetWithOnlyLC(bot *ogame.OGame, planetID ogame.PlanetID) int64 {

	eachCap := GetCapacityOfCurrentBot(bot, ogame.LargeCargo.BaseShip)
	ships, _ := bot.GetShips(planetID.Celestial())
	return ships.LargeCargo * eachCap

}
