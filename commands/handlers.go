package commands

import (
	"fmt"
	"hot-reload/animes"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

func HandleScrapingMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	c := colly.NewCollector(
		colly.AllowedDomains("https://mangayabu.top", "mangayabu.top"),
	)
	useCase := animes.NewAnimeUseCase(c)
	scrapingResult := useCase.Read()
	fmt.Println()
	if len(scrapingResult) == 0 {
		s.ChannelMessageSend(m.ChannelID, "Não foi econtrado nenhuma atualização para a data solicitada")
		return
	}

	var projectsMessage []string
	maxMessageGroupSize, _ := strconv.Atoi(os.Getenv("MAX_SITE_MESSAGE_GROUP"))
	for _, r := range scrapingResult {

		projectsMessage = append(projectsMessage, fmt.Sprintf("Project: ** %s**\nUrl: ** %s **", r.Title, r.Url))
		if len(projectsMessage) >= maxMessageGroupSize {
			message := strings.Join(projectsMessage, "\n")

			s.ChannelMessageSend(m.ChannelID, message)
			projectsMessage = nil
		}
	}
	return
}

func HandleNotFoundCommandMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Não consegui compreender a mensagem, tente mandar algo como: **!news**, que irei lhe listar todos os lançamentos que temos")
	return
}
