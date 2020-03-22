package OGameBotRoutine

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/jc01rho/ogamego_bot/OgameUtil"
	"math"
)

func (bot *OGameBot) GetNextResBuilding() (*ogame.Planet, *ogame.BaseBuilding, int64) {

	planets := bot.Ogamebot.GetCachedPlanets()
	var targetPlanet *ogame.Planet = nil
	var targetBuilding *ogame.BaseBuilding = nil
	var lowsestPrice int64 = math.MaxInt64

	var currentLevel int64 = -1

	for _, elm := range planets {

		if bot.BuildRessSkipList.Contains(elm) {
			continue
		}
		a, b, _, _ := elm.ConstructionsBeingBuilt()
		if a != 0 && b != 0 {
			continue
		}

		resbuildings, _ := elm.GetResourcesBuildings()
		energy, _ := elm.GetResources()

		temp, _ := bot.Ogamebot.GetShips(elm.GetID())

		temp2 := bot.Ogamebot.GetUserInfos()

		temp3 := bot.Ogamebot.GetServer().Settings.EspionageProbeRaids

		_ = temp2
		_ = temp
		_ = temp3

		if energy.Energy < 0 && resbuildings.SolarPlant < 22 {

			targetPlanet = &elm
			targetBuilding = &ogame.SolarPlant.BaseBuilding
			currentLevel = resbuildings.SolarPlant

			return targetPlanet, targetBuilding, currentLevel

		} else {

			if lowsestPrice > OgameUtil.ResourcePricesSum(ogame.MetalMine.GetPrice(resbuildings.MetalMine)) {
				targetPlanet = &elm
				targetBuilding = &ogame.MetalMine.BaseBuilding
				lowsestPrice = OgameUtil.ResourcePricesSum(ogame.MetalMine.GetPrice(resbuildings.MetalMine))
				currentLevel = resbuildings.MetalMine

			} else if lowsestPrice > OgameUtil.ResourcePricesSum(ogame.DeuteriumSynthesizer.GetPrice(resbuildings.DeuteriumSynthesizer)) {
				targetPlanet = &elm
				targetBuilding = &ogame.MetalMine.BaseBuilding
				lowsestPrice = OgameUtil.ResourcePricesSum(ogame.DeuteriumSynthesizer.GetPrice(resbuildings.DeuteriumSynthesizer))
				currentLevel = resbuildings.DeuteriumSynthesizer

			} else if lowsestPrice > OgameUtil.ResourcePricesSum(ogame.CrystalMine.GetPrice(resbuildings.CrystalMine)) {
				targetPlanet = &elm
				targetBuilding = &ogame.MetalMine.BaseBuilding
				lowsestPrice = OgameUtil.ResourcePricesSum(ogame.CrystalMine.GetPrice(resbuildings.CrystalMine))
				currentLevel = resbuildings.CrystalMine

			}

		}

	}

	return targetPlanet, targetBuilding, currentLevel

}
