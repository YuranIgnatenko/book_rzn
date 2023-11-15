package parsing

import (
	"encoding/csv"
	"os"

	"github.com/gocolly/colly/v2"
)

type ProsvCard struct {
	Autor string `csv:"autor"`
	Title string `csv:"title"`
	Price string `csv:"price"`
	Link  string `csv:"link"`
}

type ParsingService struct{}

func (ps *ParsingService) WriteToCsv(data []*ProsvCard) {
	file, _ := os.Open("bd_prosv.csv")
	writer := csv.NewWriter(file)
	for _, dt := range data {
		row := []string{dt.Autor, dt.Title, dt.Price, dt.Link}
		err := writer.Write(row)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()

}

// https://shop.prosv.ru/katalog?pagenumber=2
var url = "http://shop.prosv.ru/katalog"

func (ps *ParsingService) ScrapSource() []ProsvCard {
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

	c.Visit(url)

	return dts

}
