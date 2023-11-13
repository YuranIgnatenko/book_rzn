package parsing

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type ProsvCard struct {
	Autor string
	Title string
	Price string
	Link  string
}

// https://shop.prosv.ru/katalog?pagenumber=2

var url = "http://shop.prosv.ru/katalog"

func GetLinks() []ProsvCard {
	c := colly.NewCollector()
	dts := make([]ProsvCard, 0)

	c.OnHTML(".item-grid", func(e *colly.HTMLElement) {
		e.ForEach(".item-box", func(_ int, el *colly.HTMLElement) {
			dt := ProsvCard{
				Autor: el.ChildText(".autor"),
				Title: el.ChildText(".product-title"),
				Price: el.ChildText(".prices"),
				Link:  el.ChildAttr("img", "src"),
			}

			dts = append(dts, dt)
		})
	})
	fmt.Println(len(dts))

	c.Visit(url)
	return dts

}
