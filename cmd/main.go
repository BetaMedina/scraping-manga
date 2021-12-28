package main

import (
	"fmt"
	"hot-reload/core/animes"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("mangalivre.net", "https://mangalivre.net"),
	)

	useCase := animes.NewAnimeUseCase(c)
	scrapingResult := useCase.Read()
	result, err := useCase.SaveReport(scrapingResult)
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println("Operation has finished with with sucess : ", result)
}
