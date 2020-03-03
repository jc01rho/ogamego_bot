package OGameBot

import "github.com/jc01rho/ogamego_bot/OgameUtil"

func (bot *OGameBot) BuildNextRess() {

	targetPlaent, targetBuilding, _ := OgameUtil.GetNextResBuilding(bot.ogamebot)
	OgameUtil.BuildTargetBuilding(bot.ogamebot, targetPlaent.ID.Celestial(), *targetBuilding)

}
