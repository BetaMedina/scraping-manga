package main

import (
	"encoding/json"
	"fmt"
	"hot-reload/core/animes"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("mangalivre.net", "https://mangalivre.net"),
	)

	infos := animes.NewAnimeUseCase(c)
	scrapingResult := infos.Read()
	jsonResponse, _ := json.Marshal(scrapingResult)
	fmt.Println(string(jsonResponse))
}
