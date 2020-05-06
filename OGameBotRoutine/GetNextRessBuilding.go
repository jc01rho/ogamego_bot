package OGameBotRoutine

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/jc01rho/ogamego_bot/OgameUtil"
	"github.com/jc01rho/ogamego_bot/Queue"
	log "github.com/sirupsen/logrus"
	"math"
	"time"
)

func (bot *OGameBot) GetNextResBuilding() (*ogame.Planet, ogame.ID, int64) {

	planets := bot.Ogamebot.GetCachedPlanets()
	var targetPlanet ogame.Planet
	var targetBuilding ogame.Base
	var lowsestPrice int64 = math.MaxInt64

	var currentLevel int64 = -1

	for _, elm := range planets {

		if bot.BuildRessSkipList.Contains(elm) {
			continue
		}
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

			targetPlanet = elm
			targetBuilding = ogame.SolarPlant.Base
			currentLevel = resbuildings.SolarPlant

			return &targetPlanet, targetBuilding.ID, currentLevel

		} else if energy.Energy < 0  {
			targetPlanet = elm
			targetBuilding = ogame.SolarSatellite.Base
			currentLevel = 5


			go func() {
				time.Sleep(60*time.Second)
				Queue.JobQueue.Set(Queue.DefaultPriority, OGameBotGlobal.BuildNextRess)
			}()
			return &targetPlanet, targetBuilding.ID, currentLevel

		} else {

			if lowsestPrice > OgameUtil.ResourcePricesSum(ogame.MetalMine.GetPrice(resbuildings.MetalMine)) {
				targetPlanet = elm
				targetBuilding = ogame.MetalMine.Base
				lowsestPrice = OgameUtil.ResourcePricesSum(ogame.MetalMine.GetPrice(resbuildings.MetalMine))
				currentLevel = resbuildings.MetalMine

			} else if lowsestPrice > OgameUtil.ResourcePricesSum(ogame.DeuteriumSynthesizer.GetPrice(resbuildings.DeuteriumSynthesizer)) {
				targetPlanet = elm
				targetBuilding = ogame.DeuteriumSynthesizer.Base
				lowsestPrice = OgameUtil.ResourcePricesSum(ogame.DeuteriumSynthesizer.GetPrice(resbuildings.DeuteriumSynthesizer))
				currentLevel = resbuildings.DeuteriumSynthesizer

			} else if lowsestPrice > OgameUtil.ResourcePricesSum(ogame.CrystalMine.GetPrice(resbuildings.CrystalMine)) {
				targetPlanet = elm
				targetBuilding = ogame.CrystalMine.Base
				lowsestPrice = OgameUtil.ResourcePricesSum(ogame.CrystalMine.GetPrice(resbuildings.CrystalMine))
				currentLevel = resbuildings.CrystalMine

			}

		}

	}

	if &targetPlanet != nil && &targetBuilding != nil {
		log.Info("GetNextResBuilding will be return.")
		log.Infof("target ressbuilding is %s : %s %s ", targetPlanet.Coordinate, targetBuilding.Name, currentLevel+1)
	}

	return &targetPlanet, targetBuilding.ID, currentLevel

}
