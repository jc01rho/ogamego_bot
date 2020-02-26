package Test

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/jc01rho/ogamego_bot/Global"
)

func GetBotForTest() *ogame.OGame {

	//universe := os.Getenv("Vega") // eg: Bellatrix
	//username := os.Getenv("rkjnice@gmail.com") // eg: email@gmail.com
	//password := os.Getenv("aa132537") // eg: *****
	//language := os.Getenv("en") // eg: en
	universe  :="Vega"
	username  := "rkjnice@gmail.com"
	password  := "aa132537"
	language  := "en"
	bot, err := ogame.New(universe, username, password, language)

	if err != nil {
		Global.Logger.Error(err.Error())
	}

	return bot
}


