package main

import (
	"bitbucket.org/jc01rho/ogame"
	"fmt"
	"github.com/jc01rho/ogamego_bot/Telegram"
	"github.com/urfave/cli/v2"
	"github.com/utahta/go-cronowriter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}


//텔레그램 테스트 봇 토큰 / golang_test_jc / 781061948:AAHhMNvHv2RN0uIarC8DIQP9F6ShYhP0CY8
var telegramBotToken = "TOKEN"

func main() {

	w1 := cronowriter.MustNew("/tmp/example.log.%Y%m%d")
	w2 := cronowriter.MustNew("/tmp/internal_error.log.%Y%m%d")

	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder


	l := zap.New(

		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			atom,),
		zap.ErrorOutput()
	)
	l.Info("test")

	app := &cli.App{
		Action: func(c *cli.Context) error {
			fmt.Printf("Hello %q", c.Args().Get(0))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}



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
