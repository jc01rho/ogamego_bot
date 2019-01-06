package main

import (
	"ogame-golang-bot/Telegram"
)

//텔레그램 테스트 봇 토큰 / golang_test_jc / 781061948:AAHhMNvHv2RN0uIarC8DIQP9F6ShYhP0CY8
//var telegramBotToken = "781061948:AAHhMNvHv2RN0uIarC8DIQP9F6ShYhP0CY8"
var telegramBotToken = "TOKEN"

func main() {


	Telegram.BotInit(telegramBotToken)



	//universe := os.Getenv("UNIVERSE") // eg: Bellatrix
	//username := os.Getenv("USERNAME") // eg: email@gmail.com
	//password := os.Getenv("PASSWORD") // eg: *****
	//language := os.Getenv("LANGUAGE") // eg: en
	//bot, err := ogame.New(universe, username, password, language)
	//if err != nil {
	//	panic(err)
	//}
	//attacked := bot.IsUnderAttack()
	//fmt.Println(attacked) // False
}
