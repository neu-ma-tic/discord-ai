package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	//
	dg.Identify.Intents = discordgo.IntentsDirectMessages | discordgo.IntentsGuildMessages
	dg.AddHandler(messageCreate)
	dg.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{
			Name: os.Getenv("PRESENCE_TEXT"),
		},
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection:", err)
		return
	}

	fmt.Println("bot is running, ctrl-c to stop.")
	<-make(chan struct{})
	return
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	channel, _ := s.Channel(m.ChannelID)
	isDmChannel := channel.Type == discordgo.ChannelTypeDM

	if strings.HasPrefix(m.Content, fmt.Sprintf("<@%s>", s.State.User.ID)) || isDmChannel {
		args := strings.Split(m.Content, " ")
		if len(args) == 0 {
			return
		}

		// if not dm channel, omit first argument (bot ping)
		if !isDmChannel {
			args = args[1:]
		}

		blocked := strings.Split(os.Getenv("BLOCKED_USERS"), " ")
		for _, v := range blocked {
			if v == m.Author.ID {
				s.ChannelMessageSendReply(m.ChannelID, "no", m.Reference())
				return
			}
		}

		fmt.Printf("user \"%s\" called AI with prompt %s \n", m.Author.ID, args)

		res := RunAI(strings.Join(args, " "))
		s.ChannelMessageSendReply(m.ChannelID, res, m.Reference())
	}
}
