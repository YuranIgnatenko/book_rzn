package parsing

import (
	"backend/models"
	"strings"

	"github.com/gocolly/colly/v2"
)

type ServiceProsv struct {
	LinksVisit  []string
	TargetCards []models.TargetCard
	SourceType  string
	TagName     string
}

func NewServiceProsv(links []string, tagname string) *ServiceProsv {
	return &ServiceProsv{
		LinksVisit: links,
		TagName:    tagname,
		SourceType: "prosv.ru",
	}
}

func (ps *ServiceProsv) ScrapSource() []models.TargetCard {
	c := colly.NewCollector()
	dts := make([]models.TargetCard, 0)

	for _, link := range ps.LinksVisit {
		c.OnHTML(".item-grid", func(e *colly.HTMLElement) {
			e.ForEach(".item-box", func(_ int, el *colly.HTMLElement) {
				dt := models.TargetCard{
					Autor: el.ChildText(".autor"),
					Title: el.ChildText(".product-title"),
					Price: el.ChildText(".prices"),
					Link:  el.ChildAttr("img", "src"),
					Source: ps.SourceType,
					Tag: ps.TagName,
				}
				dt.Title = strings.ReplaceAll(dt.Title, `"`, "")

				dt.Id = dt.Autor + dt.Title + dt.Price + el.ChildAttr("img", "src")
				dt.TargetHash = dt.Id

				dts = append(dts, dt)
			})
		})

		c.Visit(link)
	}

	return dts

}
