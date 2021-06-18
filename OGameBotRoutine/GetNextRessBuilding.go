package OGameBotRoutine

import (
	"github.com/0xE232FE/ogame.mod"
	"github.com/jc01rho/ogamego_bot/OgameUtil"
	"github.com/jc01rho/ogamego_bot/Queue"
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"time"
)

func (bot *OGameBot) GetNextResBuilding() (*ogame.Planet, ogame.ID, int64) {
	s1 := rand.NewSource(time.Now().UnixNano())
	random := rand.New(s1)



	planets := bot.Ogamebot.GetCachedPlanets()
	var targetPlanet ogame.Planet
	var targetBuilding ogame.Base
	var lowsestPrice int64 = math.MaxInt64

	var currentLevel int64 = -1


	///행성별 반복문이니 조심,







	mainPlanet,err := bot.Ogamebot.GetPlanet(ogame.PlanetID(bot.MainCelestitial.GetID()))

	if  err != nil {
		log.Info("get planet err")
	}

	if mainPlanet.GetID() > 0{
		resbuildingMain , _ :=  mainPlanet.GetResourcesBuildings()

		//&& random.Intn(100) > 100-25
		if resbuildingMain.DeuteriumSynthesizer > 7  {
			isAvailable, resID, resLevel := bot.DoResearchForColonyShip()
			if isAvailable {




				return &mainPlanet, resID, resLevel
			}
		}
	}



	//if resbuildings.DeuteriumSynthesizer > 7 && random.Intn(100) > 75{
	//	isAvailable, resID := bot.DoResearchForColonyShip()
	//	if isAvailable {
	//
	//
	//
	//	}
	//
	//
	//}

	for _, elm := range planets {

		fields := elm.GetFields()

		//log.Info("remaing of ", fromCelestitial.GetCoordinate(), " is ", currentRess)
		log.Info(elm.GetCoordinate(), "current fields : ", fields.Built, "/", fields.Total)

		if fields.Total-fields.Built == 0 {
			log.Info("not enought field skip")
			continue
		}

		if bot.BuildRessSkipList.Contains(elm) {
			continue
		}
		a, b, _, _ := elm.ConstructionsBeingBuilt()
		if a != 0 && b != 0 {

			bot.BuildRessSkipList.Add(elm)
			continue
		} else {
			bot.BuildRessSkipList.Remove(elm)
		}

		if bot.BuildRessSkipList.Contains(elm) {
			continue
		}

		productions, _ := elm.GetResourcesProductions()
		settings, _ := elm.GetResourceSettings()

		details, _ := elm.GetResourcesDetails()
		resbuildings, _ := elm.GetResourcesBuildings()

		settings.MetalMine = 100
		settings.CrystalMine = 100
		settings.DeuteriumSynthesizer = 100
		settings.SolarPlant = 100
		settings.Crawler = 100
		settings.SolarSatellite = 100
		_ = bot.Ogamebot.SetResourceSettings(ogame.PlanetID(elm.GetID()), settings)

		//bot.Ogamebot.SetResourceSettings()
		resbuildings.CrystalMine++
		resbuildings.DeuteriumSynthesizer++
		resbuildings.MetalMine++

		_ = productions
		_ = settings
		_ = details

		if details.Energy.Available < 0 && resbuildings.SolarPlant < 22 {

			targetPlanet = elm
			targetBuilding = ogame.SolarPlant.Base
			currentLevel = resbuildings.SolarPlant

			return &targetPlanet, targetBuilding.ID, currentLevel

		} else if details.Energy.Available < 0 {

			builtList, reaminingTime, _ := elm.GetProduction()
			if len(builtList) > 0 {
				bot.BuildRessSkipList.Add(elm)
				go func() {
					//TODO : 지연 빌드 혹은 지연 큐 삽입 테스트 필요
					time.Sleep(time.Second*time.Duration(reaminingTime) + time.Second*3600)
					bot.BuildRessSkipList.Remove(targetPlanet)

				}()
				continue
			} else {
				targetPlanet = elm
				targetBuilding = ogame.SolarSatellite.Base
				currentLevel = 5
				go func() {
					time.Sleep(60 * time.Second)
					Queue.JobQueue.Set(Queue.DefaultPriority, OGameBotGlobal.BuildNextRess)
				}()
				return &targetPlanet, targetBuilding.ID, currentLevel
			}

		} else {



			if resbuildings.DeuteriumSynthesizer > 20 && random.Intn(100) > 100-30 {
				result, tempTargetPlanet, tempTargetBuildingID, tempCurrentLevel := bot.GetNextFacilityInBuildRess(elm)
				if result {
					return &tempTargetPlanet, tempTargetBuildingID, tempCurrentLevel
				}

			} else if resbuildings.DeuteriumSynthesizer > 16 && random.Intn(100) > 100-25 {
				result, tempTargetPlanet, tempTargetBuildingID, tempCurrentLevel := bot.GetNextFacilityInBuildRess(elm)
				if result {
					return &tempTargetPlanet, tempTargetBuildingID, tempCurrentLevel
				}

			} else if resbuildings.DeuteriumSynthesizer > 12 && random.Intn(100) > 100-20 {
				result, tempTargetPlanet, tempTargetBuildingID, tempCurrentLevel := bot.GetNextFacilityInBuildRess(elm)
				if result {
					return &tempTargetPlanet, tempTargetBuildingID, tempCurrentLevel
				}

			}




			if details.Metal.StorageCapacity < productions.Metal*32 {
				targetPlanet = elm
				targetBuilding = ogame.MetalStorage.Base
				lowsestPrice = OgameUtil.ResourcePricesSum(ogame.MetalStorage.GetPrice(resbuildings.MetalStorage))
			} else if details.Crystal.StorageCapacity < productions.Crystal*32 {
				targetPlanet = elm
				targetBuilding = ogame.CrystalStorage.Base
				lowsestPrice = OgameUtil.ResourcePricesSum(ogame.CrystalStorage.GetPrice(resbuildings.CrystalStorage))
			} else if details.Deuterium.StorageCapacity < productions.Deuterium*32 {
				targetPlanet = elm
				targetBuilding = ogame.DeuteriumTank.Base
				lowsestPrice = OgameUtil.ResourcePricesSum(ogame.DeuteriumTank.GetPrice(resbuildings.DeuteriumTank))
			}

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

	return &targetPlanet, targetBuilding.ID, currentLevel

}
