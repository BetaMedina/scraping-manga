package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ExecuteCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	message := strings.Split(strings.TrimSpace(m.Content), " ")

	switch message[0] {
	case "!news":
		fmt.Println(message)
		HandleScrapingMessages(s, m)
		break

	default:
		HandleNotFoundCommandMessage(s, m)
		break
	}
}
