package parsing

import (
	"backend/config"
	"backend/connector"
	"backend/models"
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func TargetHash(autor, title, price, image string) string {
	return fmt.Sprintf("%v%v%v%v", autor, title, price, image)
}

type ParsingService struct {
	LinkVisit   []string
	TargetCards []models.TargetCard
	config.Configuration
	connector.Connector
}

func NewParsingService(c config.Configuration, conn connector.Connector) *ParsingService {
	ps := ParsingService{
		Configuration: c,
		Connector:     conn,
		LinkVisit: []string{
			"http://shop.prosv.ru/katalog",
			"https://shop.prosv.ru/katalog?pagenumber=2",
		},
	}

	data := conn.GetListTargets()
	ps.TargetCards = make([]models.TargetCard, 0)
	if len(data) <= 1 {
		ps.TargetCards = ps.ScrapSource()
		for _, card := range ps.TargetCards {
			target_hash := TargetHash(card.Autor, card.Title, card.Price, card.Link)
			conn.SaveTargetTargets(target_hash, card.Autor, card.Title, card.Price, card.Link)
		}
		return &ps

	} else {
		ps.TargetCards = data
	}

	return &ps
}

func (ps *ParsingService) ScrapSource() []models.TargetCard {
	c := colly.NewCollector()
	dts := make([]models.TargetCard, 0)

	for _, link := range ps.LinkVisit {
		c.OnHTML(".item-grid", func(e *colly.HTMLElement) {
			e.ForEach(".item-box", func(_ int, el *colly.HTMLElement) {
				dt := models.TargetCard{
					Autor: el.ChildText(".autor"),
					Title: el.ChildText(".product-title"),
					Price: el.ChildText(".prices"),
					Link:  el.ChildAttr("img", "src"),
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
