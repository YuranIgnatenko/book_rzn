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
	LinkVisitProsv                    []string
	LinkVisitAgatFreshNewTable        []string
	LinkVisitAgatFreshNewBasicModules []string
	LinkVisitAgatStudentTable         []string
	LinkVisitAgatStudentChair         []string
	LinkVisitAgatOfficeOptimaTable    []string
	LinkVisitAgatOfficeOptimaModules  []string

	TargetCardsProsv []models.TargetCard
	TargetCardsAgat  []models.TargetCard

	config.Configuration
	connector.Connector
}

func NewParsingService(c config.Configuration, conn connector.Connector) *ParsingService {
	ps := ParsingService{
		Configuration: c,
		Connector:     conn,
		LinkVisitProsv: []string{
			"http://shop.prosv.ru/katalog",
			"https://shop.prosv.ru/katalog?pagenumber=2",
		},
		LinkVisitAgatFreshNewTable: []string{
			"https://agatmk.ru/stolyi-rabochie-fresh#menustart",
		},
		LinkVisitAgatFreshNewBasicModules: []string{
			"https://agatmk.ru/moduli-sistemyi-xraneniya-fresh#menustart",
		},
		LinkVisitAgatStudentTable: []string{
			"https://agatmk.ru/uchenicheskie-stolyi#menustart",
		},
		LinkVisitAgatStudentChair: []string{
			"https://agatmk.ru/uchenicheskie-stulya#menustart",
		},
		LinkVisitAgatOfficeOptimaTable: []string{
			"https://agatmk.ru/stolyi-rabochie-optima#menustart",
		},
		LinkVisitAgatOfficeOptimaModules: []string{
			"https://agatmk.ru/sistema-xraneniya-optima#menustart",
		},
	}

	// data := conn.GetListTargets()
	data := []models.TargetCard{}

	ps.TargetCardsProsv = make([]models.TargetCard, 0)
	ps.TargetCardsAgat = make([]models.TargetCard, 0)

	if len(data) <= 1 {
		// ps.TargetCardsProsv = ps.ScrapSourceProsv()
		ps.TargetCardsProsv = []models.TargetCard{}
		ps.TargetCardsAgat = ps.ScrapSourceAgatOfficeOptimaModules()

		for _, el := range ps.TargetCardsAgat {
			fmt.Printf("%#+v\n\n", el)
		}

		for _, card := range ps.TargetCardsProsv {
			target_hash := TargetHash(card.Autor, card.Title, card.Price, card.Link)
			conn.SaveTargetTargets(target_hash, card.Autor, card.Title, card.Price, card.Link)
		}
		return &ps

	} else {
		ps.TargetCardsProsv = data
	}

	return &ps
}
