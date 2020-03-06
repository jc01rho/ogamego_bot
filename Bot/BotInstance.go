package Bot

import (
	"bitbucket.org/jc01rho/ogame"
	"github.com/jc01rho/ogamego_bot/OGameBot"
	"os"
)

func InitBot() {
	universe := os.Getenv("UNIVERSE") // eg: Bellatrix
	username := os.Getenv("USERNAME") // eg: email@gmail.com
	password := os.Getenv("PASSWORD") // eg: *****
	language := os.Getenv("LANGUAGE") // eg: en
	bot, _ := ogame.New(universe, username, password, language)
	OGameBot.OGameBotGlobal.Ogamebot = bot
}
