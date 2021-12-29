package commands

import (
	"strconv"
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
		page, _ := strconv.Atoi(message[1])
		HandleScrapingMessages(s, m, page)
		break

	default:
		HandleNotFoundCommandMessage(s, m)
		break
	}
}
