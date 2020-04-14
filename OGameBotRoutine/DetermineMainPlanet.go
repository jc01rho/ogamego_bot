package OGameBotRoutine

import (
	"bitbucket.org/jc01rho/ogame"
)

func (bot *OGameBot) DetermineMainPlanet() {

	planets := bot.Ogamebot.GetCachedPlanets()
	var lowsestID *ogame.Planet = nil
	var coordness ogame.Coordinate
	//TODO : 달을 메인행성으로 지정하기.
	//TODO : 달이 메인행성이 된다면, 자원량 확인하여 카고 충원 및 자원 모으기를 동시에 해야함.
	//최소한도로 전행성 생산량의24시간 분을 달에서 운송할수있게 하고, 향후 전계정자원량 +10%를 가질수있도록?
	for _, elm := range planets {
		if lowsestID != nil {
			if lowsestID.GetID() > elm.GetID() {
				lowsestID = &elm
				coordness = elm.GetCoordinate()
			}

		} else {
			lowsestID = &elm
			coordness = elm.GetCoordinate()
		}
	}

	bot.MainPlanetCoord = coordness
	bot.MainPlanetCelestitial = lowsestID
	if lowsestID != nil && lowsestID.Moon != nil {
		bot.MainPlanetMoonCelestitial = lowsestID.Moon
	} else {
		bot.MainPlanetMoonCelestitial = nil
	}

	bot.IsMainPlanetMoon = false

}
