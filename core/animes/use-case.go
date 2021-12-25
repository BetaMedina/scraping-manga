package animes

import (
	"strconv"

	"github.com/gocolly/colly"
)

type Scraping interface {
	Read() *AnimeScraping
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
	for i := 1; i < 6; i++ {
		s.c.Visit("https://mangalivre.net/lista-de-mangas/ordenar-por-atualizacoes?page=" + strconv.Itoa(i))
	}
	return infos
}
