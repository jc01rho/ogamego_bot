package OGameBot

import (
	"github.com/jc01rho/ogamego_bot/Logger"
	"github.com/jc01rho/ogamego_bot/OgameUtil"
)

func (bot *OGameBot) BuildNextRess() {

	Logger.Logger.Info("BuildNextRess start")
	targetPlaent, targetBuilding, level := OgameUtil.GetNextResBuilding(bot.Ogamebot)
	CurrentRessInTargetPlanet, _ := bot.Ogamebot.GetResources(targetPlaent.ID.Celestial())
	NeedsRess := targetBuilding.GetPrice(level + 1)
	bot.Ogamebot.Abandon()

	if CurrentRessInTargetPlanet.CanAfford(NeedsRess) {
		//자원 부족시 모아서 처리하는 기능 구현 필요
		//달 위주 운영 하자...

	} else {
		OgameUtil.BuildTargetBuilding(bot.Ogamebot, targetPlaent.ID.Celestial(), *targetBuilding)
	}

}
