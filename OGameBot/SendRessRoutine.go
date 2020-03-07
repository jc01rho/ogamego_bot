package OGameBot

import "bitbucket.org/jc01rho/ogame"

func (bot *OGameBot) SendRess(from *ogame.Planet, amount ogame.Resources, to *ogame.Planet) {

	currentRess, _ := bot.Ogamebot.GetResources(from.ID.Celestial())

	remaining := currentRess.SubReal(amount)
	if remaining.Metal < 0 || remaining.Crystal < 0 || remaining.Deuterium < 0 {

	}

	//bot.Ogamebot.SendFleet()

}
