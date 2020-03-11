package OGameBot

import (
	"github.com/jc01rho/ogamego_bot/OgameUtil"
	log "github.com/sirupsen/logrus"
)

func (bot *OGameBot) BuildNextRess() {

	log.Info("BuildNextRess start")
	targetPlaent, targetBuilding, level := OgameUtil.GetNextResBuilding(bot.Ogamebot)
	CurrentRessInTargetPlanet, _ := bot.Ogamebot.GetResources(targetPlaent.ID.Celestial())
	NeedsRess := targetBuilding.GetPrice(level + 1)
	//bot.Ogamebot.Abandon()

	if CurrentRessInTargetPlanet.CanAfford(NeedsRess) {
		log.Error("Not Enough Ress for buildRess")
		//자원 부족시 모아서 처리하는 기능 구현 필요
		//달 위주 운영 하자...

	} else {
		OgameUtil.BuildTargetBuilding(bot.Ogamebot, targetPlaent.ID.Celestial(), *targetBuilding)
	}

}
