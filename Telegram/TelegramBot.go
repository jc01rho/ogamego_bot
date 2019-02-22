package Telegram

import (

	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func init () {

}

func BotInit(botToken string) {
	b, err := tb.NewBot(tb.Settings{
		Token:  botToken,
		// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
		//URL: "http://195.129.111.17:8012",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

	b.Start()
}