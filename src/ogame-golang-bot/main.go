package main

import (
	"fmt"
	"github.com/alaingilbert/ogame"
	"github.com/marcsantiago/gocron"
	"ogame-golang-bot/Telegram"
)

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}


//텔레그램 테스트 봇 토큰 / golang_test_jc / 781061948:AAHhMNvHv2RN0uIarC8DIQP9F6ShYhP0CY8
//var telegramBotToken = "781061948:AAHhMNvHv2RN0uIarC8DIQP9F6ShYhP0CY8"
var telegramBotToken = "781061948:AAHhMNvHv2RN0uIarC8DIQP9F6ShYhP0CY8"
//type PriorityQueue []*Item
func main() {

	gocron.Every(1).Second().Do(taskWithParams, 1, "hello")
	gocron.Every(2).Second().Do(taskWithParams, 2, "hello")

	 go gocron.Start()


	go Telegram.BotInit(telegramBotToken)




	fmt.Println("Start OGBot Routine")


	universe := "Dorado" // eg: Bellatrix
	username := "rkjnice@gmail.com" // eg: email@gmail.com
	password := "aa132537" // eg: *****
	language := "en" // eg: en
	bot, err := ogame.New(universe, username, password, language)
	if err != nil {
		panic(err)
	}
	attacked := bot.IsUnderAttack()
	fmt.Println(attacked) // False
}
