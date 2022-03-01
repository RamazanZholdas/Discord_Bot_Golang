package bot

import (
	"discordTestBot/config"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	BotId string
	goBot *discordgo.Session
	err   error
)

func Start() {
	goBot, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
	}

	u, err := goBot.User("@me")
	if err != nil {
		log.Fatal(err)
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		log.Fatal(err)
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	} else if !strings.HasPrefix(m.Content, config.BotPrefix) {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
