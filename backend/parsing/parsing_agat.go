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

func (ps *ParsingService) ScrapSourceAgatFreshNewTables() []models.TargetCard {
	c := colly.NewCollector()
	dts := make([]models.TargetCard, 0)

	for _, link := range ps.LinkVisitAgatFreshNewTable {
		c.OnHTML(".cats_list_container", func(e *colly.HTMLElement) {
			e.ForEach(".product", func(_ int, el *colly.HTMLElement) {
				dt := models.TargetCard{
					Autor: "NO Autor",
					Title: el.ChildText(".goods_list_name"),
					Price: "NO price",
					Link:  "https://agatmk.ru" + el.ChildAttr("img", "src"),
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

func (ps *ParsingService) ScrapSourceAgatFreshNewBasicModules() []models.TargetCard {
	c := colly.NewCollector()
	dts := make([]models.TargetCard, 0)

	for _, link := range ps.LinkVisitAgatFreshNewBasicModules {
		c.OnHTML(".cats_list_container", func(e *colly.HTMLElement) {
			e.ForEach(".product", func(_ int, el *colly.HTMLElement) {
				dt := models.TargetCard{
					Autor: "NO Autor",
					Title: el.ChildText(".goods_list_name"),
					Price: "NO price",
					Link:  "https://agatmk.ru" + el.ChildAttr("img", "src"),
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

func (ps *ParsingService) ScrapSourceAgatStudentTable() []models.TargetCard {
	c := colly.NewCollector()
	dts := make([]models.TargetCard, 0)

	for _, link := range ps.LinkVisitAgatStudentTable {
		c.OnHTML(".cats_list_container", func(e *colly.HTMLElement) {
			e.ForEach(".product", func(_ int, el *colly.HTMLElement) {
				dt := models.TargetCard{
					Autor: "NO Autor",
					Title: el.ChildText(".goods_list_name"),
					Price: "NO price",
					Link:  "https://agatmk.ru" + el.ChildAttr("img", "src"),
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

func (ps *ParsingService) ScrapSourceAgatStudentChair() []models.TargetCard {
	c := colly.NewCollector()
	dts := make([]models.TargetCard, 0)

	for _, link := range ps.LinkVisitAgatStudentChair {
		c.OnHTML(".cats_list_container", func(e *colly.HTMLElement) {
			e.ForEach(".product", func(_ int, el *colly.HTMLElement) {
				dt := models.TargetCard{
					Autor: "NO Autor",
					Title: el.ChildText(".goods_list_name"),
					Price: "NO price",
					Link:  "https://agatmk.ru" + el.ChildAttr("img", "src"),
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

func (ps *ParsingService) ScrapSourceAgatOfficeOptimaTable() []models.TargetCard {
	c := colly.NewCollector()
	dts := make([]models.TargetCard, 0)

	for _, link := range ps.LinkVisitAgatOfficeOptimaTable {
		c.OnHTML(".cats_list_container", func(e *colly.HTMLElement) {
			e.ForEach(".product", func(_ int, el *colly.HTMLElement) {
				dt := models.TargetCard{
					Autor: "NO Autor",
					Title: el.ChildText(".goods_list_name"),
					Price: "NO price",
					Link:  "https://agatmk.ru" + el.ChildAttr("img", "src"),
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

func (ps *ParsingService) ScrapSourceAgatOfficeOptimaModules() []models.TargetCard {
	c := colly.NewCollector()
	dts := make([]models.TargetCard, 0)

	for _, link := range ps.LinkVisitAgatOfficeOptimaModules {
		c.OnHTML(".cats_list_container", func(e *colly.HTMLElement) {
			e.ForEach(".product", func(_ int, el *colly.HTMLElement) {
				dt := models.TargetCard{
					Autor: "NO Autor",
					Title: el.ChildText(".goods_list_name"),
					Price: "NO price",
					Link:  "https://agatmk.ru" + el.ChildAttr("img", "src"),
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
