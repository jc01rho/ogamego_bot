package OGameBot

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/jc01rho/ogamego_bot/OgameUtil"
	log "github.com/sirupsen/logrus"
	"math"
)

func (bot *OGameBot) MaintainLCCountStep(LCcount int64) {


	planets := bot.Ogamebot.GetCachedPlanets()

	for _, elm := range planets {
		bot.MaintainLCCountStepSpecificPlanet(elm,LCcount)
	}


}


//LCCount가 -1 이면 시간당 생산량에서 카대를 생산
func (bot *OGameBot) MaintainLCCountStepSpecificPlanet(planet ogame.Planet, LCcount int64) {
	log.Info("MaintainLCCountStep start")

	elm := planet


		Resss, _ := elm.GetResourcesProductions() //TODO : this function fecthes resss percent to crawl. consider that check minlevel and calc for production.

		currentCap := OgameUtil.GetCapacityOfCurrentBotOfPlanetWithOnlyLC(bot.Ogamebot, elm.ID)
		ToBuildLCCount := int64(1)
		if Resss.Total()*24 > currentCap { //handle to take 24h productions
			//TODO: build LC more
			if LCcount < 0 {
				ToBuildLCCount = int64(math.Min(float64(Resss.Metal/ogame.LargeCargo.Price.Metal), float64(Resss.Crystal/ogame.LargeCargo.Price.Crystal)))
				ToBuildLCCount = int64(math.Max(float64(ToBuildLCCount), 1))
			} else {
				ToBuildLCCount = (LCcount)
			}

			log.Info("%v Build LC %v", elm.Coordinate, ToBuildLCCount)

			//Logger.Logger.Debug("BuildLC",zap.Object)
			elm.BuildShips(ogame.LargeCargo.ID, int64(ToBuildLCCount))
		}

	log.Info("MaintainLCCountStep done")

}
