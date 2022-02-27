package main

import (
	"discordTestBot/bot"
	"discordTestBot/config"
	"log"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot.Start()

	<-make(chan struct{})
}
