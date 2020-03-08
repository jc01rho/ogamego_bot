package OGameBot

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/jc01rho/ogamego_bot/OgameUtil"
	"math"
)

func (bot *OGameBot) SendRess(from *ogame.Planet, amount ogame.Resources, to *ogame.Planet) {

	currentRess, _ := bot.Ogamebot.GetResources(from.ID.Celestial())
	currentResearch := bot.Ogamebot.GetCachedResearch()

	remaining := currentRess.SubReal(amount)
	if remaining.Metal < 0 || remaining.Crystal < 0 || remaining.Deuterium < 0 {
		//TODO: not enough ress
		//return
	}

	currentShips, _ := bot.Ogamebot.GetShips(from.ID.Celestial())
	LCCaps := ogame.LargeCargo.GetCargoCapacity(currentResearch, bot.Ogamebot.GetServer().Settings.EspionageProbeRaids != 0, bot.Class == ogame.Collector)
	//SCCaps := ogame.SmallCargo.GetCargoCapacity(currentResearch, bot.Ogamebot.GetServer().Settings.EspionageProbeRaids != 0 , bot.Class == ogame.Collector)
	SCCaps := OgameUtil.GetCapacityOfCurrentBot(bot.Ogamebot, ogame.SmallCargo.BaseShip)

	currentCaps := LCCaps*currentShips.LargeCargo + SCCaps + currentShips.SmallCargo

	if OgameUtil.ResourcePricesSum(amount) > currentCaps {
		//TODO: not enouth cargo
		//return
	}

	////real sending part
	//only LC available,
	if LCCaps*currentShips.LargeCargo >= OgameUtil.ResourcePricesSum(amount) {
		NeededLCs := int64(math.Ceil(float64(OgameUtil.ResourcePricesSum(amount)) / float64(LCCaps)))
		_, Fuels := bot.Ogamebot.FlightTime(from.Coordinate, to.Coordinate, ogame.HundredPercent, ogame.ShipsInfos{LargeCargo: NeededLCs})
		NeededLCs += int64(math.Ceil(float64(Fuels) / float64(LCCaps)))

		if NeededLCs >= currentShips.LargeCargo {
			ships := make([]ogame.Quantifiable, 0)
			ships = append(ships, ogame.Quantifiable{ID: ogame.LightFighterID, Nbr: int64(NeededLCs)})

			bot.Ogamebot.SendFleet(from.ID.Celestial(), ships, ogame.HundredPercent, to.Coordinate, ogame.Transport, amount, 0, 0)

		}

	}

	//bot.Ogamebot.SendFleet()

}
