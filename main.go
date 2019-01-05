package main

import "fmt"
import "os"
import "github.com/alaingilbert/ogame"

func main() {
	universe := os.Getenv("UNIVERSE") // eg: Bellatrix
	username := os.Getenv("USERNAME") // eg: email@gmail.com
	password := os.Getenv("PASSWORD") // eg: *****
	language := os.Getenv("LANGUAGE") // eg: en
	bot, err := ogame.New(universe, username, password, language)
	if err != nil {
		panic(err)
	}
	attacked := bot.IsUnderAttack()
	fmt.Println(attacked) // False
}