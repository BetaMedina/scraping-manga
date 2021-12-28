package main

import (
	"fmt"
	"hot-reload/animes"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	message := strings.Split(m.Content, " ")

	if message[0] == "!news" {
		c := colly.NewCollector(
			colly.AllowedDomains("mangalivre.net", "https://mangalivre.net"),
		)
		useCase := animes.NewAnimeUseCase(c)

		page, _ := strconv.Atoi(message[1])
		scrapingResult := useCase.Read(page)

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
	s.ChannelMessageSend(m.ChannelID, "Não consegui compreender a mensagem, tente mandar algo como: !news Pagina, que irei lhe dizer os lançamentos que temos")
	return
}
