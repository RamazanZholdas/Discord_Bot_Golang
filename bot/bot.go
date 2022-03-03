package bot

import (
	"discordTestBot/config"
	"discordTestBot/functions"
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
	}
	request := m.Content[4:]
	if !strings.HasPrefix(m.Content, config.BotPrefix) {
		return
	}

	for _, v := range request {
		if v > 123 {
			request = functions.Translating(request)
			break
		}
	}

	slice := functions.ParseData(request)

	if slice != nil {
		for i := range slice {
			s.ChannelMessageSend(m.ChannelID, slice[i])
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "movie doesnt exist")
	}
}
