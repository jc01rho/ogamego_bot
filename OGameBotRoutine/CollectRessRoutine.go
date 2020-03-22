package OGameBotRoutine

import (
	"github.com/jc01rho/ogamego_bot/OgameUtil"
	log "github.com/sirupsen/logrus"
)

func (bot *OGameBot) CollectRessRoutine() {

	planets := bot.Ogamebot.GetCachedPlanets()
	for _, elm := range planets {

		if elm.Coordinate != bot.MainPlanetCoord {
			log.Info(elm.Coordinate, bot.MainPlanetCoord, "is diff")
			currentRess, _ := elm.GetResources()
			currentCap := OgameUtil.GetCapacityOfCurrentBotOfPlanetWithOnlyLC(bot.Ogamebot, elm.ID)

			if currentRess.Total() >= currentCap {

				bot.MaintainLCCountStepSpecificPlanet(elm, -1)

			} else {
				bot.SendRessRoutineFromCelestitial(elm.ID.Celestial(), currentRess, bot.MainPlanetCoord)
			}

		}
	}
}
