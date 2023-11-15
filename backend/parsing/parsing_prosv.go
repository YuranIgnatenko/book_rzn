package parsing

import (
	"backend/bd"
	"backend/config"
	"backend/models"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

type ParsingService struct {
	LinkVisit      []string
	ProsvCardCache []models.ProsvCard
	config.Configuration
	bd.Bd
}

func NewParsingService(c config.Configuration, bd bd.Bd) *ParsingService {
	ps := ParsingService{
		Configuration: c,
		Bd:            bd,
		LinkVisit: []string{
			"http://shop.prosv.ru/katalog",
			"https://shop.prosv.ru/katalog?pagenumber=2",
		},
	}
	data, _ := ps.ReadFromCsv()
	// fmt.Println(data, err)
	fmt.Println("++++")
	if len(data) <= 1 {
		ps.ProsvCardCache = ps.ScrapSource()
		ps.WriteToCsv(ps.ProsvCardCache)
		return &ps

	} else {
		// data = ps.ScrapSource()
		fmt.Println(len(data))
		ps.ProsvCardCache = data
		// ps.WriteToCsv(data)
	}

	return &ps
}

func (ps *ParsingService) WriteToCsv(data []models.ProsvCard) {
	file, err := os.OpenFile(ps.Path_bd+ps.Bd_prosv, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, dt := range data {
		row := []string{dt.Autor, dt.Title, dt.Price, dt.Link}
		err := writer.Write(row)
		if err != nil {
			panic(err)
		}
	}

}

func (ps *ParsingService) ScrapSource() []models.ProsvCard {
	c := colly.NewCollector()
	dts := make([]models.ProsvCard, 0)

	for _, link := range ps.LinkVisit {
		c.OnHTML(".item-grid", func(e *colly.HTMLElement) {
			e.ForEach(".item-box", func(_ int, el *colly.HTMLElement) {
				dt := models.ProsvCard{
					Autor: el.ChildText(".autor"),
					Title: el.ChildText(".product-title"),
					Price: el.ChildText(".prices"),
					Link:  el.ChildAttr("img", "src"),
				}

				dts = append(dts, dt)
			})
		})

		c.Visit(link)
	}

	return dts

}

func (ps *ParsingService) ReadFromCsv() ([]models.ProsvCard, error) {
	data := []models.ProsvCard{}

	file, err := os.Open(ps.Path_bd + ps.Bd_prosv)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 0
	reader.Comment = '#'

	for {
		rec, err := reader.Read()
		if err != nil {
			if len(data) >= 0 {
				err = nil
				return data, err
			} else {
				return nil, err
			}
		}
		fmt.Println(rec)
		card := models.ProsvCard{
			Autor: rec[0],
			Title: rec[1],
			Price: rec[2],
			Link:  rec[3],
		}
		data = append(data, card)
	}
	// if len(data) < 5 {
	// 	return nil, errors.New("BD empty, not found records")
	// }
	return data, err

}
