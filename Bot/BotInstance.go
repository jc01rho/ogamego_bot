package Bot

import (
	"bitbucket.org/jc01rho/ogame"
	"os"
)

var BotInstance *ogame.OGame

func InitBot() {
	universe := os.Getenv("UNIVERSE") // eg: Bellatrix
	username := os.Getenv("USERNAME") // eg: email@gmail.com
	password := os.Getenv("PASSWORD") // eg: *****
	language := os.Getenv("LANGUAGE") // eg: en
	bot, _ := ogame.New(universe, username, password, language)
	BotInstance = bot
}
