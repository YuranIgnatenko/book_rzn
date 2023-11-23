// https://agatmk.ru/moduli-sistemyi-xraneniya-fresh#menustart

package parsing

import (
	"backend/models"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gocolly/colly/v2"
)

func htmlToString(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

type ServiceAgat struct {
	LinksVisit  []string
	TargetCards []models.TargetCard
	SourceType  string
	TagName     string
}

func NewServiceAgat(links []string, tagname string) *ServiceAgat {
	return &ServiceAgat{
		LinksVisit: links,
		TagName:    tagname,
		SourceType: "agatmk.ru",
	}
}

func (sa *ServiceAgat) ScrapSource() []models.TargetCard {
	c := colly.NewCollector()
	dts := make([]models.TargetCard, 0)

	for _, link := range sa.LinksVisit {
		c.OnHTML(".cats_list_container", func(e *colly.HTMLElement) {
			e.ForEach(".product", func(_ int, el *colly.HTMLElement) {
				dt := models.TargetCard{
					Autor:  "NO Autor",
					Title:  el.ChildText(".goods_list_name"),
					Price:  "NO price",
					Link:   "https://agatmk.ru" + el.ChildAttr("img", "src"),
					Source: sa.SourceType,
					Tag:    sa.TagName,
				}
				dt.Title = strings.ReplaceAll(dt.Title, `"`, "")
				dt.Id = dt.Title + dt.Link
				dt.TargetHash = dt.Id
				dts = append(dts, dt)
			})
		})
		c.Visit(link)
	}
	return dts
}
