package main

import (
	"github.com/jc01rho/ogamego_bot/Global"
	"github.com/jc01rho/ogamego_bot/Logger"
)

//텔레그램 테스트 봇 토큰 / golang_test_jc / 781061948:AAHhMNvHv2RN0uIarC8DIQP9F6ShYhP0CY8
var telegramBotToken = "TOKEN"

func main() {

	Logger.InitLogger()

	//go Telegram.BotInit(telegramBotToken)

	Global.InitEssentials()

	blocking := make(chan int)

	<-blocking

}
