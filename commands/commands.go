package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func ExecuteCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	message := strings.Split(strings.TrimSpace(m.Content), " ")
	switch message[0] {
	case "!news":
		HandleScrapingMessages(s, m)
		break

	case "!chapters":
		chapter := strings.Join(removeIndex(message, 0), "-")
		HandleScrapingFindChapters(s, m, chapter)
		break

	default:
		HandleNotFoundCommandMessage(s, m)
		break
	}
}
