package OGameBotRoutine

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/jc01rho/ogamego_bot/OgameUtil"
	log "github.com/sirupsen/logrus"
	"math"
)

func (bot *OGameBot) SendRessRoutineFromCelestitial(fromCelestitial ogame.CelestialID, amount ogame.Resources, to ogame.Coordinate) bool {
	currentRess, _ := bot.Ogamebot.GetResources(fromCelestitial)
	currentResearch := bot.Ogamebot.GetCachedResearch()

	remaining := currentRess.SubReal(amount)
	if remaining.Metal < 0 || remaining.Crystal < 0 || remaining.Deuterium < 0 {
		log.Error("NotEnough Ress for send")
		//TODO: not enough ress
		return false
	}

	currentShips, _ := bot.Ogamebot.GetShips(fromCelestitial)
	LCCaps := ogame.LargeCargo.GetCargoCapacity(currentResearch, bot.Ogamebot.GetServer().Settings.EspionageProbeRaids != 0, bot.Class == ogame.Collector)
	//SCCaps := ogame.SmallCargo.GetCargoCapacity(currentResearch, bot.Ogamebot.GetServer().Settings.EspionageProbeRaids != 0 , bot.Class == ogame.Collector)
	SCCaps := OgameUtil.GetCapacityOfCurrentBot(bot.Ogamebot, ogame.SmallCargo.BaseShip)

	currentCaps := LCCaps*currentShips.LargeCargo + SCCaps + currentShips.SmallCargo

	if OgameUtil.ResourcePricesSum(amount) > currentCaps {
		//TODO: not enouth cargo
		return false
	}

	////real sending part
	//only LC available,
	if LCCaps*currentShips.LargeCargo >= OgameUtil.ResourcePricesSum(amount) {
		NeededLCs := int64(math.Ceil(float64(OgameUtil.ResourcePricesSum(amount)) / float64(LCCaps)))
		_, Fuels := bot.Ogamebot.FlightTime(bot.Ogamebot.GetCachedCelestialByID(fromCelestitial).GetCoordinate(), to, ogame.HundredPercent, ogame.ShipsInfos{LargeCargo: NeededLCs})
		NeededLCs += int64(math.Ceil(float64(Fuels) / float64(LCCaps)))

		if NeededLCs <= currentShips.LargeCargo {
			ships := make([]ogame.Quantifiable, 0)
			ships = append(ships, ogame.Quantifiable{ID: ogame.LargeCargoID, Nbr: int64(NeededLCs)})

			bot.Ogamebot.SendFleet(fromCelestitial, ships, ogame.HundredPercent, to, ogame.Transport, amount, 0, 0)

			return true
		}

	}
	return false
}

func (bot *OGameBot) SendRessRoutine(from *ogame.Planet, amount ogame.Resources, to ogame.Coordinate) bool {
	return bot.SendRessRoutineFromCelestitial(from.ID.Celestial(), amount, to)
}

func (bot *OGameBot) SendRessRoutineTargetPlanet(from *ogame.Planet, amount ogame.Resources, to *ogame.Planet) bool {
	return bot.SendRessRoutineFromCelestitial(from.ID.Celestial(), amount, to.Coordinate)
}