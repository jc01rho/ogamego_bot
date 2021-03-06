package OGameBotRoutine

import "github.com/0xE232FE/ogame.mod"

func (bot *OGameBot) GetNextFacilityInBuildRess(planet ogame.Planet) (bool, ogame.Planet, ogame.ID, int64) {

	//targetBuilding = ogame.MetalMine.Base
	//lowsestPrice = OgameUtil.ResourcePricesSum(ogame.MetalMine.GetPrice(resbuildings.MetalMine))
	//currentLevel = resbuildings.MetalMine
	facilities, _ := planet.GetFacilities()
	researches := bot.Ogamebot.GetResearch()

	if facilities.RoboticsFactory < 2 {
		return true, planet, ogame.RoboticsFactoryID, facilities.RoboticsFactory
	} else if facilities.Shipyard < 2 {
		return true, planet, ogame.ShipyardID, facilities.Shipyard
	} else if facilities.Shipyard < 4 {
		return true, planet, ogame.ShipyardID, facilities.Shipyard
	} else if facilities.ResearchLab < 4 {
		return true, planet, ogame.ShipyardID, facilities.ResearchLab
	} else if facilities.RoboticsFactory < 7 {
		return true, planet, ogame.RoboticsFactoryID, facilities.RoboticsFactory
	} else if facilities.Shipyard < 8 {
		return true, planet, ogame.ShipyardID, facilities.Shipyard
	} else if facilities.ResearchLab < 8 {
		return true, planet, ogame.ShipyardID, facilities.ResearchLab
	} else if facilities.RoboticsFactory < 10 {
		return true, planet, ogame.RoboticsFactoryID, facilities.RoboticsFactory
	} else if facilities.ResearchLab < 10 {
		return true, planet, ogame.ShipyardID, facilities.ResearchLab
	} else if facilities.NaniteFactory < 2 && researches.ComputerTechnology >= 10 {
		return true, planet, ogame.RoboticsFactoryID, facilities.RoboticsFactory
	} else if facilities.ResearchLab < 12 {
		return true, planet, ogame.ShipyardID, facilities.ResearchLab
	} else {
		//fail fall
		return false, planet, ogame.ShipyardID, facilities.Shipyard
	}

}
