package OGameBotRoutine

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/jc01rho/ogamego_bot/OgameUtil"
)

func (bot *OGameBot) CollectRessRoutine() {

	planets := bot.Ogamebot.GetCachedPlanets()
	for _, elm := range planets {

		if elm.Coordinate != bot.MainPlanetCoord {
			//log.Info(elm.Coordinate, bot.MainPlanetCoord, "is diff")
			currentRess, _ := elm.GetResources()
			currentCap := OgameUtil.GetCapacityOfCurrentBotOfPlanetWithOnlyLC(bot.Ogamebot, elm.ID)

			if currentRess.Total() >= currentCap {

				bot.MaintainLCCountStepSpecificPlanet(elm, -1)

				//크리를 운송량의 절반가까이 채움
				currentRess, _ := elm.GetResources()
				var toSend ogame.Resources = struct {
					Metal      int64
					Crystal    int64
					Deuterium  int64
					Energy     int64
					Darkmatter int64
				}{Metal: 0, Crystal: 0, Deuterium: 0, Energy: 0, Darkmatter: 0}

				if float64(currentRess.Crystal) >= float64(currentCap)*0.5 {
					toSend.Crystal = int64(float64(currentCap) * 0.5)
				} else {
					toSend.Crystal = currentRess.Crystal
				}

				toSend.Metal = currentRess.Metal

				if toSend.Total() >= currentCap {
					toSend.Metal -= (toSend.Total() - currentCap)
				}

				if toSend.Total() == currentCap {
					bot.SendRessRoutineFromCelestitial(elm, toSend, bot.MainPlanetCoord)
				} else {
					toSend.Deuterium = currentRess.Deuterium
					if toSend.Total() >= currentCap {
						toSend.Deuterium -= (toSend.Total() - currentCap)
					}
					bot.SendRessRoutineFromCelestitial(elm, toSend, bot.MainPlanetCoord)
				}

			} else {
				bot.SendRessRoutineFromCelestitial(elm, currentRess, bot.MainPlanetCoord)
			}

		}
	}
}
