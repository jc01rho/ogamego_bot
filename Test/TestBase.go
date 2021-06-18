package Test

import (
	"github.com/0xE232FE/ogame.mod"
)

func GetBotForTest() *ogame.OGame {

	//Global.InitLogger()

	//universe := os.Getenv("Vega") // eg: Bellatrix
	//username := os.Getenv("rkjnice@gmail.com") // eg: email@gmail.com
	//password := os.Getenv("aa132537") // eg: *****
	//language := os.Getenv("en") // eg: en
	universe := "Vega"
	username := "rkjnice@gmail.com"
	password := "aa132537"
	language := "en"
	bot, _ := ogame.New(universe, username, password, language)

	return bot
}
