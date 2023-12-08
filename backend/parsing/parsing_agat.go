// https://agatmk.ru/moduli-sistemyi-xraneniya-fresh#menustart

// https://agatmk.ru/moduli-sistemyi-xraneniya-fresh#menustart

package parsing

import (
	"backend/models"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

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
				price := strings.Split(el.Text, "Цена")[1]
				price = strings.Split(price, ":")[1]
				price = strings.Split(price, ".")[0]
				price = strings.TrimSpace(price)

				if strings.Contains(price, "0 руб") || strings.Contains(price, "по запросу") {
					return
					// e.DOM.Next() ??
				}

				temp_price := strings.ReplaceAll(price, "\u00a0", "")
				temp_price = strings.ReplaceAll(temp_price, " ", "")
				temp_price = strings.ReplaceAll(temp_price, "руб", "")
				temp_price = strings.TrimSpace(temp_price)

				dt := models.TargetCard{
					Autor:  "-",
					Title:  el.ChildText(".goods_list_name"),
					Price:  temp_price,
					Link:   "https://agatmk.ru" + el.ChildAttr("img", "src"),
					Source: sa.SourceType,
					Tag:    sa.TagName,
				}

				dt.Title = strings.ReplaceAll(dt.Title, `"`, "")
				dt.Id = dt.Autor + dt.Title + dt.Price + "https://agatmk.ru" + el.ChildAttr("img", "src")

				dt.TargetHash = fmt.Sprintf("%v", time.Now().UnixNano())

				dts = append(dts, dt)
			})
		})
		c.Visit(link)
	}
	return dts
}

