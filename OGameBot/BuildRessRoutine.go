package OGameBot

import (
	"github.com/jc01rho/ogamego_bot/Logger"
	"github.com/jc01rho/ogamego_bot/OgameUtil"
)

func (bot *OGameBot) BuildNextRess() {

	Logger.Logger.Info("BuildNextRess start")
	targetPlaent, targetBuilding, _ := OgameUtil.GetNextResBuilding(bot.Ogamebot)
	OgameUtil.BuildTargetBuilding(bot.Ogamebot, targetPlaent.ID.Celestial(), *targetBuilding)

}
