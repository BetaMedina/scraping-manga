package animes

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

type Scraping interface {
	Read() *AnimeScraping
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

	s.c.OnHTML("div#titulos-az > div.content-wraper > div.tag-container> div.seriesList", func(e *colly.HTMLElement) {
		e.ForEach("ul.seriesList > li", func(_ int, el *colly.HTMLElement) {
			formattedValues := AnimeScraping{
				Url:   `https://mangalivre.net` + el.ChildAttrs("li > a.link-block ", "href")[0],
				Title: el.ChildText("li > a.link-block > div.series-info > span.series-title > h1 "),
			}
			infos = append(infos, &formattedValues)
		})
	})
	for i := 1; i < 13; i++ {
		s.c.Visit("https://mangalivre.net/lista-de-mangas/ordenar-por-atualizacoes?page=" + strconv.Itoa(i))
	}
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
