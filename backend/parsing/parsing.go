package parsing

import (
	"backend/config"
	"backend/connector"
	"backend/models"
	"fmt"
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
		ps.TargetCards = ps.ScrapSourceProsv()
		// ps.TargetCards = ps.ScrapSourceAgat()

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
