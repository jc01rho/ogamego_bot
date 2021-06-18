package OGameBotRoutine

import (
	"github.com/0xE232FE/ogame.mod"
	"github.com/jc01rho/ogamego_bot/Logger"
	"github.com/jc01rho/ogamego_bot/OgameUtil"
	"github.com/jc01rho/ogamego_bot/Queue"
	log "github.com/sirupsen/logrus"
	"time"
)

func (bot *OGameBot) BuildNextRess() {

	log.Info("BuildNextRess start")
	targetPlanet, targetID, level := bot.GetNextResBuilding()
	CurrentRessInTargetPlanet, _ := bot.Ogamebot.GetResources(targetPlanet.GetID())
	targetObject := ogame.Objs.ByID(targetID)
	var NeedsRess ogame.Resources

	//if ogame.Objs.ByID(targetID) != nil && ogame.Objs.ByID(targetID).GetID().IsResourceBuilding() {
	//
	//	NeedsRess = targetObject.GetPrice(level)
	//} else if ogame.Objs.ByID(targetID) != nil && ogame.Objs.ByID(targetID).GetID().IsShip() && targetID == ogame.SolarSatelliteID {
	//	NeedsRess = targetObject.GetPrice(level)
	//} else {
	//	log.Error(targetID, "Wrong targetID ID ")
	//	//return
	//}

	NeedsRess = targetObject.GetPrice(level)

	//bot.Ogamebot.Abandon()

	if bot.BuildRessSkipList.Contains(targetPlanet) {
		log.Info("")
		return
	}

	bot.BuildRessSkipList.Add(targetPlanet)

	log.Infof("Will build %s %s %s", targetPlanet.Coordinate.String(), targetObject.GetName(), string(level+1))
	log.Infof("Ress Needed :  %s", NeedsRess.String())
	log.Info("CurrentRess in Planet : %s ", CurrentRessInTargetPlanet.String())

	if !CurrentRessInTargetPlanet.CanAfford(NeedsRess) {
		log.Info("Not Enough Ress for buildRess, Checking Main")
		RessinMainPlanet, _ := bot.Ogamebot.GetResources(bot.MainPlanetCelestitial.GetID())
		if CurrentRessInTargetPlanet.Add(RessinMainPlanet).CanAfford(NeedsRess) {

			log.Infof("%s Sending ress from Main[%s] to %s", Logger.CurrentFileNameAndLine(), bot.MainPlanetCoord, targetPlanet.Coordinate)
			isSuccesful := bot.SendRessRoutineFromCelestitial(bot.MainPlanetCelestitial, NeedsRess.Sub(CurrentRessInTargetPlanet), targetPlanet.Coordinate)
			if isSuccesful {
				flightTime, _ := bot.Ogamebot.FlightTime(bot.MainPlanetCelestitial.GetCoordinate(), targetPlanet.Coordinate, ogame.HundredPercent, ogame.ShipsInfos{LargeCargo: 1})

				go func() {
					//TODO : 지연 빌드 혹은 지연 큐 삽입 테스트 필요
					log.Infof("%s Sleep %d secs and Build command will be added to queue", Logger.CurrentFileNameAndLine(), flightTime+30)
					time.Sleep(time.Second*time.Duration(flightTime) + time.Second*30)
					log.Infof(">>>>RessBuilding Lazy Built %s %s", targetPlanet.Coordinate.String(), targetObject.GetName())
					Queue.JobQueue.Set(Queue.DefaultPriority, func() { OgameUtil.BuildTargetID(bot.Ogamebot, targetPlanet.GetID(), targetObject, level) })
					bot.BuildRessSkipList.Remove(targetPlanet)

				}()
			}


		} else {
			log.Info("Not Enough Ress Altough sumed with main, abort build")
			bot.BuildRessSkipList.Remove(targetPlanet)
		}

		//자원 부족시 모아서 처리하는 기능 구현 필요
		//달 위주 운영 하자...

	} else {

		OgameUtil.BuildTargetID(bot.Ogamebot, targetPlanet.GetID(), targetObject, level)
		log.Info("RessBuilding Immdiately Built")
		bot.BuildRessSkipList.Remove(targetPlanet)
	}

}
