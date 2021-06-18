package OGameBotRoutine

import (
	"github.com/0xE232FE/ogame.mod"
	log "github.com/sirupsen/logrus"
)

func (bot *OGameBot) BuildDefs() {

	log.Info("BuildDefs Routine Start")

	planets := bot.Ogamebot.GetCachedPlanets()
	for _, elm := range planets {
		productions, _, errors := elm.GetProduction()
		if errors != nil {
			continue
		}

		if len(productions) > 7 {
			log.Info("productions leng over7, skipped")
			continue
		} else {
			log.Info("build defs inner routine start...")
			resbuildings, err := elm.GetResourcesBuildings()
			if err != nil {
				continue
			}

			if resbuildings.DeuteriumSynthesizer > 15 {
				_ = elm.BuildDefense(ogame.GaussCannonID, 2)
			}

			if resbuildings.DeuteriumSynthesizer > 20 {
				_ = elm.BuildDefense(ogame.PlasmaTurretID, 1)
			}

			if resbuildings.CrystalMine > 10 {
				_ = elm.BuildDefense(ogame.LightLaserID, 5)
			}

			if resbuildings.CrystalMine > 20 {
				_ = elm.BuildDefense(ogame.LightLaserID, 4)
			}

			if resbuildings.MetalMine > 15 {
				_ = elm.BuildDefense(ogame.RocketLauncherID, 3)
			}
			if resbuildings.MetalMine > 20 {
				_ = elm.BuildDefense(ogame.RocketLauncherID, 5)

			}

		}

	}

	log.Info("BuildDefs Routine Done")

}
