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
			fmt.Println("++++::::", target_hash)
			conn.SaveTargetTargets(target_hash, card.Autor, card.Title, card.Price, card.Link)
		}
		// ps.WriteToCsv(ps.ProsvCards)
		return &ps

	} else {
		ps.TargetCards = data
	}

	return &ps
}

// func (ps *ParsingService) WriteToCsv(data []models.ProsvCard) {
// 	file, err := os.OpenFile(ps.Path_bd+ps.Bd_prosv, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

// 	if err != nil {
// 		panic(err)
// 	}
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()
// 	for _, dt := range data {
// 		row := []string{dt.Autor, dt.Title, dt.Price, dt.Link, dt.Id}
// 		err := writer.Write(row)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// }

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
				// dt.Title = strings.ReplaceAll(dt.Title, "'", "")
				// dt.Title = strings.TrimSpace(dt.Title)
				// dt.Title = strings.ReplaceAll(dt.Title, "//", "/")

				// dt.Link = strings.ReplaceAll(dt.Link, "\"", "")
				// dt.Link = strings.ReplaceAll(dt.Link, "'", "")
				// dt.Link = strings.TrimSpace(dt.Link)
				// dt.Link = strings.ReplaceAll(dt.Link, "/", " ")

				dt.Id = dt.Autor + dt.Title + dt.Price + el.ChildAttr("img", "src")
				dt.TargetHash = dt.Id

				// dt.Title = strings.ReplaceAll(dt.Title, "\\", "")
				fmt.Println("parsinnnng:::::::", dt.Title)
				dts = append(dts, dt)
			})
		})

		c.Visit(link)
	}

	return dts

}

// func (ps *ParsingService) ReadFromCsv() ([]models.ProsvCard, error) {
// 	data := []models.ProsvCard{}

// 	file, err := os.Open(ps.Path_bd + ps.Bd_prosv)

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	reader.FieldsPerRecord = 0
// 	reader.Comment = '#'

// 	for {
// 		rec, err := reader.Read()
// 		if err != nil {
// 			if len(data) >= 0 {
// 				err = nil
// 				return data, err
// 			} else {
// 				return nil, err
// 			}
// 		}
// 		fmt.Println(rec)
// 		card := models.ProsvCard{
// 			Autor: rec[0],
// 			Title: rec[1],
// 			Price: rec[2],
// 			Link:  rec[3],
// 			Id:    rec[4],
// 		}
// 		data = append(data, card)
// 	}
// }
