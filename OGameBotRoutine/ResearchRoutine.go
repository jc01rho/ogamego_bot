package OGameBotRoutine

import "bitbucket.org/jc01rho/ogame"

func (bot *OGameBot) DoResearchForColonyShip() (bool, ogame.ID, int64) {

	researches := bot.Ogamebot.GetCachedResearch()
	researchesFinishat := bot.Ogamebot.GetResearchFinishAt()

	if researchesFinishat > 0 {
		return false, -1, 0
	}

	print("test")
	_ = researches
	_ = researchesFinishat

	if researches.EnergyTechnology < 1 {
		return true, ogame.EnergyTechnologyID, researches.EnergyTechnology

	} else if researches.CombustionDrive < 2 {
		return true, ogame.CombustionDriveID, researches.CombustionDrive

	} else if researches.EspionageTechnology < 4 {
		return true, ogame.EspionageTechnologyID, researches.EspionageTechnology

	} else if researches.CombustionDrive < 3 {
		return true, ogame.CombustionDriveID, researches.CombustionDrive

	} else if researches.ImpulseDrive < 3 {
		return true, ogame.ImpulseDriveID, researches.ImpulseDrive

	} else if researches.Astrophysics < 1 {
		return true, ogame.AstrophysicsID, researches.Astrophysics

	}

	//if facilities['research_lab'] < 3:
	//self.logger.info("build ResearchLab start")
	//self.OGame.build(pidLocal, Facilities['ResearchLab'])
	//
	//
	//elif self.researchList['energy_technology'] < 1:
	//self.logger.info("research energy_technology from  " + str(self.researchList['energy_technology']))
	//self.OGame.build(self.Settings['collect']['main']["target"]['id'], Research['EnergyTechnology'])
	//
	//elif self.researchList['combustion_drive'] < 2:
	//self.logger.info("research combustion_drive from  " + str(self.researchList['combustion_drive']))
	//self.OGame.build(self.Settings['collect']['main']['target']['id'], Research['CombustionDrive'])
	//
	//elif self.researchList['espionage_technology'] < 4:
	//self.logger.info("research espionage_technology from  " + str(self.researchList['espionage_technology']))
	//self.OGame.build(self.Settings['collect']['main']['target']['id'], Research['EspionageTechnology'])
	//
	//elif self.researchList['combustion_drive'] < 3:
	//self.logger.info("research combustion_drive from  " + str(self.researchList['combustion_drive']))
	//self.OGame.build(self.Settings['collect']['main']['target']['id'], Research['CombustionDrive'])
	//
	//elif facilities['research_lab'] < 4:
	//self.logger.info("build ResearchLab start")
	//self.OGame.build(pidLocal, Facilities['ResearchLab'])
	//
	//elif self.researchList['impulse_drive'] < 3:
	//self.logger.info("research impulse_drive from  " + str(self.researchList['impulse_drive']))
	//self.OGame.build(self.Settings['collect']['main']['target']['id'], Research['ImpulseDrive'])
	//
	//
	//elif facilities['shipyard'] < 5:
	//self.logger.info("build shipyard start")
	//self.OGame.build(pidLocal, Facilities['shipyard'])
	//
	//elif self.researchList['astrophysics'] < 1:
	//self.logger.info("research astrophysics from  " + str(self.researchList['astrophysics']))
	//self.OGame.build(self.Settings['collect']['main']['target']['id'], Research['Astrophysics'])

	return false, -1, 0
}

func (bot *OGameBot) DoResearch() {

}
