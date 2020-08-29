package OGameBotRoutine

import (
	"bitbucket.org/jc01rho/ogame"
	log "github.com/sirupsen/logrus"
)

func (bot *OGameBot) BuildDefs() {

	log.Info("BuildDefs Routine Start")

	planets := bot.Ogamebot.GetCachedPlanets()
	for _, elm := range planets {
		productions, _, _ := elm.GetProduction()
		if len(productions) > 5 {
			continue
		} else {

			resbuildings, _ := elm.GetResourcesBuildings()
			if resbuildings.MetalMine > 15 {
				_ = elm.BuildDefense(ogame.RocketLauncherID, 2)
			}
			if resbuildings.MetalMine > 20 {
				_ = elm.BuildDefense(ogame.RocketLauncherID, 3)

				if resbuildings.CrystalMine > 10 {
					_ = elm.BuildDefense(ogame.LightLaserID, 1)
				}

				if resbuildings.CrystalMine > 20 {
					_ = elm.BuildDefense(ogame.LightLaserID, 3)
				}

				if resbuildings.DeuteriumSynthesizer > 15 {
					_ = elm.BuildDefense(ogame.GaussCannonID, 2)
				}

				if resbuildings.DeuteriumSynthesizer > 20 {
					_ = elm.BuildDefense(ogame.PlasmaTurretID, 1)
				}

			}

		}

	}

	log.Info("BuildDefs Routine Done")

}
