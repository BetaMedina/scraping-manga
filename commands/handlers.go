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

func HandleScrapingMessages(s *discordgo.Session, m *discordgo.MessageCreate, page int) {
	c := colly.NewCollector(
		colly.AllowedDomains("mangalivre.net", "https://mangalivre.net"),
	)
	useCase := animes.NewAnimeUseCase(c)
	scrapingResult := useCase.Read(page)
	fmt.Println(page)
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
	s.ChannelMessageSend(m.ChannelID, "Não consegui compreender a mensagem, tente mandar algo como: !news Pagina, que irei lhe dizer os lançamentos que temos")
	return
}
