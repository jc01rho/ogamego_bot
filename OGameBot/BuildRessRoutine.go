package OGameBot

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/jc01rho/ogamego_bot/Logger"
	"github.com/jc01rho/ogamego_bot/OgameUtil"
	"github.com/jc01rho/ogamego_bot/Queue"
	log "github.com/sirupsen/logrus"
	"time"
)

func (bot *OGameBot) BuildNextRess() {

	log.Info("BuildNextRess start")
	targetPlaent, targetBuilding, level := OgameUtil.GetNextResBuilding(bot.Ogamebot)
	CurrentRessInTargetPlanet, _ := bot.Ogamebot.GetResources(targetPlaent.ID.Celestial())
	NeedsRess := targetBuilding.GetPrice(level + 1)
	//bot.Ogamebot.Abandon()

	if !CurrentRessInTargetPlanet.CanAfford(NeedsRess) {
		log.Info("Not Enough Ress for buildRess, Checking Main")
		RessinMainPlanet, _ := bot.Ogamebot.GetResources(bot.MainPlanetCelestitial)
		if CurrentRessInTargetPlanet.Add(RessinMainPlanet).CanAfford(NeedsRess) {

			log.Infof("%s Sending ress from Main[%s] to %s", Logger.CurrentFileNameAndLine(), bot.MainPlanetCoord, targetPlaent.Coordinate)
			bot.SendRessRoutineFromCelestitial(bot.MainPlanetCelestitial, NeedsRess.Sub(CurrentRessInTargetPlanet), targetPlaent.Coordinate)
			timnes, _ := bot.Ogamebot.FlightTime(bot.Ogamebot.GetCachedCelestialByID(bot.MainPlanetCelestitial).GetCoordinate(), targetPlaent.Coordinate, ogame.HundredPercent, ogame.ShipsInfos{LargeCargo: 1})

			go func() {
				//TODO : 지연 빌드 혹은 지연 큐 삽입 테스트 필요
				log.Infof("%s Sleep %d secs and Build command will be added to queue", Logger.CurrentFileNameAndLine(), timnes+30)
				time.Sleep(time.Second*time.Duration(timnes) + time.Second*30)
				Queue.JobQueue.Set(func() { OgameUtil.BuildTargetBuilding(bot.Ogamebot, targetPlaent.ID.Celestial(), *targetBuilding) })

			}()

		} else {
			log.Info("Not Enough Ress Altough sumed with main, abort build")
		}

		//자원 부족시 모아서 처리하는 기능 구현 필요
		//달 위주 운영 하자...

	} else {
		OgameUtil.BuildTargetBuilding(bot.Ogamebot, targetPlaent.ID.Celestial(), *targetBuilding)
	}

}
