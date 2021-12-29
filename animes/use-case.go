package animes

import (
	"strings"

	"github.com/gocolly/colly"
)

const limitFindMangaQuantity = 5

type Scraping interface {
	Read() []*AnimeScraping
	FindManga(manga string, listQuantity int) []*ChapterScraping
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

func (s *UseCase) FindManga(manga string, listQuantity int) []*ChapterScraping {
	var infos []*ChapterScraping
	s.c.OnHTML("div#maingo > div.row > div.manga-info ", func(e *colly.HTMLElement) {
		e.ForEach("div.manga-chapters > div.single-chapter", func(_ int, el *colly.HTMLElement) {
			if len(infos) <= limitFindMangaQuantity {
				date := strings.TrimSpace(el.ChildText("small"))
				title := strings.TrimSpace(el.ChildAttr("a", "alt"))
				url := strings.TrimSpace(el.ChildAttr("a", "href"))
				formattedValues := ChapterScraping{
					Url:   url,
					Title: title,
					Date:  date,
				}
				infos = append(infos, &formattedValues)
			}
		})
	})
	s.c.Visit("https://mangayabu.top/manga/" + manga)
	return infos
}
