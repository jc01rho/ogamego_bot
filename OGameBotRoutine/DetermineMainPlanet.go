package OGameBotRoutine

import (
	"github.com/0xE232FE/ogame.mod"
	log "github.com/sirupsen/logrus"
)

func (bot *OGameBot) DetermineMainPlanet() {

	if OGameBotGlobal.MainCoords == nil {
		planets := bot.Ogamebot.GetCachedPlanets()
		var lowsestID *ogame.Planet = nil
		var coordness ogame.Coordinate
		//TODO : 달을 메인행성으로 지정하기.
		//TODO : 달이 메인행성이 된다면, 자원량 확인하여 카고 충원 및 자원 모으기를 동시에 해야함.
		//최소한도로 전행성 생산량의24시간 분을 달에서 운송할수있게 하고, 향후 전계정자원량 +10%를 가질수있도록?
		for _, elm := range planets {
			if lowsestID != nil {
				if lowsestID.GetID() > elm.GetID() {
					tempElm := elm
					lowsestID = &tempElm
					coordness = tempElm.GetCoordinate()
				}

			} else {
				tempElm := elm
				lowsestID = &tempElm
				coordness = tempElm.GetCoordinate()
			}
		}
		bot.MainPlanetCoord = coordness
		log.Info("Default selection logic  answers : ", bot.MainCoords)
		//bot.MainPlanetCelestitial = lowsestID
		bot.SetMainCelestial(lowsestID)
		if lowsestID != nil && lowsestID.Moon != nil {
			//TODO: 달이 있다면?
			bot.SetMainCelestial(lowsestID)
		} else {

			bot.SetMainCelestial(lowsestID)
		}

		bot.IsMainPlanetMoon = false

	} else {
		bot.MainPlanetCoord = *OGameBotGlobal.MainCoords

		planets := bot.Ogamebot.GetCachedPlanets()
		isExist := false
		for _, elm := range planets {
			tempElm := elm
			if tempElm.GetCoordinate().Equal(bot.MainPlanetCoord) {
				bot.SetMainCelestial(tempElm)
				isExist = true


			}
			if bot.MainPlanetCoord.IsMoon() {
				bot.IsMainPlanetMoon = true
				bot.SetMainCelestial(tempElm.Moon)
				isExist = true
			}

		}

		if isExist == false {
			OGameBotGlobal.MainCoords = nil
			log.Info("Given default coords is Null, setting to default selection logic")
			bot.DetermineMainPlanet()

		}

	}
	//log.Info(elm.Coordinate, bot.MainPlanetCoord, "is diff")
	log.Info("Main coords is", bot.MainPlanetCoord)

}
