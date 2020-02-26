package OgameUtil

import "bitbucket.org/jc01rho/ogame"

func GetNextResBuilding(bot ogame.OGame ) {

	planets := bot.GetPlanets()

	for _,elm := range planets {
		resbuildings,_ := elm.GetResourcesBuildings()
		energy,_ := elm.GetResources()

		ogame.CrystalMine.GetPrice()
		bot.GetResources()
		ogame.getR

	}


}