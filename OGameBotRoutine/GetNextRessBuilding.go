package OGameBotRoutine

import (
	"bitbucket.org/jc01rho/ogame"

	"github.com/jc01rho/ogamego_bot/OgameUtil"
	log "github.com/sirupsen/logrus"
	"math"
)

func (bot *OGameBot) GetNextResBuilding() (*ogame.Planet, *ogame.BaseBuilding, int64) {

	log.Info("GetNextResBuilding start")
	planets := bot.Ogamebot.GetCachedPlanets()
	var targetPlanet *ogame.Planet = nil
	var targetBuilding *ogame.BaseBuilding = nil
	var lowsestPrice int64 = math.MaxInt64

	var currentLevel int64 = -1

	for _, elm := range planets {

		a, b, _, _ := elm.ConstructionsBeingBuilt()
		if a != 0 && b != 0 {
			log.Info("GetNextResBuilding start")
			continue
		} else {
			bot.BuildRessSkipList.Remove(elm)
		}

		if bot.BuildRessSkipList.Contains(elm) {
			continue
		}

		resbuildings, _ := elm.GetResourcesBuildings()
		resbuildings.CrystalMine++
		resbuildings.DeuteriumSynthesizer++
		resbuildings.MetalMine++
		energy, _ := elm.GetResources()

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

	if targetPlanet != nil && targetBuilding != nil {
		log.Info("GetNextResBuilding will be return.")
		log.Infof("target ressbuilding is %s : %s %s ", targetPlanet.Coordinate, targetBuilding.Name, currentLevel+1)
	}

	return targetPlanet, targetBuilding, currentLevel

}
