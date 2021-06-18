package OGameBotRoutine

import (
//"github.com/0xE232FE/ogame.mod"
//"github.com/jc01rho/ogamego_bot/OgameUtil"
//
//"time"
)

var minArrivalTime int64 = 300
var maxArrivalTime int64 = 900

func (bot *OGameBot) DoFleetSave() {
	//var refreshAttacks int64 = 0
	//
	//var random int64 = OgameUtil.RandRange(minArrivalTime, maxArrivalTime)
	//var Saver = map[int64]ogame.Fleet{}
	//for {
	//	AutoFleetSaverOnAttackEnabledMu.RLock()
	//	Enabled := AutoFleetSaverOnAttackEnabled
	//	AutoFleetSaverOnAttackEnabledMu.RUnlock()
	//	if Enabled {
	//
	//		if refreshAttacks < time.Now().Unix() {
	//			random = randRange(minArrivalTime, maxArrivalTime)
	//			refreshAttacks = time.Now().Unix() + random
	//			bot.GetAttacks()
	//			log.Printf("[AutoFleetsaverOnAttack.go] Getting new Attacks at %s \n", time.Unix(refreshAttacks, 0).String())
	//		}
	//
	//		cachedData := bot.GetCachedData()
	//
	//		attacks := cachedData.AttackEvents
	//		if len(attacks) > 0 {
	//			for _, attack := range attacks {
	//				if attack.ArrivalTime.Unix()-600 < time.Now().Unix() {
	//					if Saver[attack.AttackerID].ID != 0 {
	//						var attacker ogame.Attacker
	//						attacker.Add(*attack.Ships)
	//						var defender ogame.Defender
	//						defender.Add(cachedData.PlanetShipsInfos[bot.GetCachedCelestialByCoord(attack.Destination).GetID()])
	//						defender.DefensesInfos = cachedData.PlanetDefensesInfos[bot.GetCachedCelestialByCoord(attack.Destination).GetID()]
	//						var params ogame.SimulatorParams
	//						params.Simulations = 100
	//						result := ogame.Simulate(attacker, defender, params)
	//						log.Printf("[AutoFleetsaverOnAttack.go] Combat Reults AttackerWin %d vs %d DefenderWin\n", result.AttackerWin, result.DefenderWin)
	//						if result.AttackerWin > result.DefenderWin {
	//							var fb ogame.FleetBuilder
	//							fb.SetAllShips()
	//							fb.SetAllResources()
	//							fb.SetRecallIn(random)
	//							fb.SetDestination(getDestination(bot, attack.Destination))
	//
	//							if cachedData.PlanetShipsInfos[bot.GetCachedCelestialByCoord(attack.Destination).GetID()].HasMovableShips() {
	//								log.Printf("[AutoFleetsaverOnAttack.go] try to Fleetsave %s \n", cachedData.PlanetDefensesInfos[bot.GetCachedCelestialByCoord(attack.Destination).GetID()].String())
	//								fleet, err := fb.SendNow()
	//								if err != nil {
	//									log.Printf("[AutoFleetsaverOnAttack.go] Error on Fleetsend %s\n", attack.Origin.String())
	//								} else {
	//									log.Printf("[AutoFleetsaverOnAttack.go] Succesfully send from %s to %s \n", attack.Origin.String(), fleet.Destination.String())
	//									Saver[attack.AttackerID] = fleet
	//								}
	//							} else {
	//								log.Printf("[AutoFleetsaverOnAttack.go] No Ships to save on Planet %s \n", attack.Origin.String())
	//							}
	//						}
	//					}
	//				} else {
	//					log.Printf("[AutoFleetsaverOnAttack.go] Attack on %s (%s) from %s (%s) is arriving at %s \n", attack.Destination, attack.DestinationName, attack.AttackerName, attack.Origin.String(), attack.ArrivalTime.String())
	//				}
	//			}
	//		}
	//	}
	//	time.Sleep(1 * time.Second)
	//}
}
