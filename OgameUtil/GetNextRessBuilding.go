package OgameUtil

import (
	"bitbucket.org/jc01rho/ogame"
	"math"
)

func GetNextResBuilding(bot *ogame.OGame )  (*ogame.Planet,*ogame.BaseBuilding,int64 ){

	planets := bot.GetPlanets()
	var targetPlanet *ogame.Planet = nil
	var targetBuilding *ogame.BaseBuilding = nil
	var lowsestPrice int64 =  math.MaxInt64
	var currentLevel int64 = -1

	for _,elm := range planets {
		resbuildings,_ := elm.GetResourcesBuildings()
		energy,_ := elm.GetResources()


		if energy.Energy < 0  && resbuildings.SolarPlant < 22{


			targetPlanet = &elm
			targetBuilding = &ogame.SolarPlant.BaseBuilding
			currentLevel =  resbuildings.SolarPlant

			return targetPlanet,targetBuilding,currentLevel


		} else {

			if lowsestPrice > ResourcePricesSum(ogame.MetalMine.GetPrice(resbuildings.MetalMine)) {
				targetPlanet = &elm
				targetBuilding = &ogame.MetalMine.BaseBuilding
				lowsestPrice = ResourcePricesSum(ogame.MetalMine.GetPrice(resbuildings.MetalMine))
				currentLevel =  resbuildings.MetalMine

			}  else if lowsestPrice > ResourcePricesSum(ogame.DeuteriumSynthesizer.GetPrice(resbuildings.DeuteriumSynthesizer)) {
				targetPlanet = &elm
				targetBuilding = &ogame.MetalMine.BaseBuilding
				lowsestPrice = ResourcePricesSum(ogame.DeuteriumSynthesizer.GetPrice(resbuildings.DeuteriumSynthesizer))
				currentLevel =  resbuildings.DeuteriumSynthesizer

			}  else if lowsestPrice > ResourcePricesSum(ogame.CrystalMine.GetPrice(resbuildings.CrystalMine)) {
				targetPlanet = &elm
				targetBuilding = &ogame.MetalMine.BaseBuilding
				lowsestPrice = ResourcePricesSum(ogame.CrystalMine.GetPrice(resbuildings.CrystalMine))
				currentLevel =  resbuildings.CrystalMine

			}




		}




	}

	return  targetPlanet,targetBuilding,currentLevel


}