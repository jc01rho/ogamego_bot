package main

import (
	"fmt"
	"bitbucket.org/jc01rho/ogame"
	"github.com/jc01rho/gocron"
	"ogame-golang-bot/Telegram"
	"os"
)

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}


//텔레그램 테스트 봇 토큰 / golang_test_jc / 781061948:AAHhMNvHv2RN0uIarC8DIQP9F6ShYhP0CY8
var telegramBotToken = "TOKEN"

func main() {

	//gocron.Every(1).Second().Do(taskWithParams, 1, "hello")
	//gocron.Every(2).Second().Do(taskWithParams, 2, "hello")

	//go gocron.Start()


	go Telegram.BotInit(telegramBotToken)




	fmt.Println("Start OGBot Routine")


	universe := os.Getenv("UNIVERSE") // eg: Bellatrix
	username := os.Getenv("USERNAME") // eg: email@gmail.com
	password := os.Getenv("PASSWORD") // eg: *****
	language := os.Getenv("LANGUAGE") // eg: en
	bot, err := ogame.New(universe, username, password, language)
	if err != nil {
		panic(err)
	}
	attacked, err := bot.IsUnderAttack()
	if err != nil {
		panic(err)
	}
	fmt.Println(attacked) // False
}
