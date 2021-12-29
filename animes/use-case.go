package animes

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Scraping interface {
	Read(newsPage int) *AnimeScraping
	SaveReport(info *AnimeScraping) (bool, error)
}

type UseCase struct {
	c *colly.Collector
}

func NewAnimeUseCase(colector *colly.Collector) *UseCase {
	return &UseCase{
		c: colector,
	}
}

func (s *UseCase) Read() []*AnimeScraping {
	var infos []*AnimeScraping
	s.c.OnHTML("div#maingo > div.row", func(e *colly.HTMLElement) {
		e.ForEach("div.s6", func(_ int, el *colly.HTMLElement) {
			url := strings.TrimSpace(el.ChildAttr("div.card > div.card-image > a", "href"))
			title := strings.TrimSpace(el.ChildText("div.card > div.card-content"))
			if title != "" && url != "" {
				formattedValues := AnimeScraping{
					Url:   url,
					Title: title,
				}
				infos = append(infos, &formattedValues)
			}
		})
	})
	s.c.Visit("https://mangayabu.top")
	return infos
}

func (s *UseCase) SaveReport(info []*AnimeScraping) (bool, error) {
	f, err := os.Create("updates.csv")
	defer f.Close()

	if err != nil {
		return false, err
	}

	w := csv.NewWriter(f)

	records := [][]string{
		{"title", "url"},
	}

	for _, k := range info {
		records = append(records, []string{
			k.Title, k.Url,
		})
	}
	err = w.WriteAll(records)
	if err != nil {
		return false, err
	}
	return true, nil
}
