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
	ListServices        []models.ServiceScraper
	ListTargetCardCache []models.TargetCard

	config.Configuration
	connector.Connector
}

func NewParsingService(c config.Configuration, conn connector.Connector) *ParsingService {
	lss := []models.ServiceScraper{
		NewServiceProsv([]string{"http://shop.prosv.ru/katalog", "https://shop.prosv.ru/katalog?pagenumber=2"}, "prosv-book"),
		NewServiceAgat([]string{"https://agatmk.ru/moduli-sistemyi-xraneniya-fresh#menustart"}, "school-boxing"),
		NewServiceAgat([]string{"https://agatmk.ru/stolyi-rabochie-fresh#menustart"}, "school-fresh-table"),
		NewServiceAgat([]string{"https://agatmk.ru/sistema-xraneniya#menustart"}, "school-fresh-boxing"),
		NewServiceAgat([]string{"https://agatmk.ru/uchenicheskie-stolyi#menustart"}, "school-table"),
		NewServiceAgat([]string{"https://agatmk.ru/uchenicheskie-stulya#menustart"}, "school-chair"),
		NewServiceAgat([]string{"https://agatmk.ru/stolyi-rabochie-optima#menustart"}, "office-optima-table"),
		NewServiceAgat([]string{"https://agatmk.ru/uchenicheskie-stulya#menustart"}, "office-optima-chair"),
		NewServiceAgat([]string{"https://agatmk.ru/tumbyi-pod-dosku#menustart"}, "school-box-from-table"),
	}

	ps := ParsingService{
		Configuration: c,
		Connector:     conn,
		ListServices:  lss,
	}

	tc_all := conn.GetListTargets()
	fmt.Println(len(tc_all))
	// tc_all := []models.TargetCard{}

	if len(tc_all) <= 1 {
		tc_all = RangeScrapServices(lss)

		for _, tc_temp := range tc_all {
			conn.SaveParsingService(tc_temp)
		}
		ps.ListTargetCardCache = tc_all

	} else {
		ps.ListTargetCardCache = tc_all
	}

	fmt.Println(len(tc_all), len(ps.ListTargetCardCache))
	return &ps
}

func RangeScrapServices(data []models.ServiceScraper) []models.TargetCard {
	tc := make([]models.TargetCard, 0)

	for _, service := range data {
		tc = append(tc, service.ScrapSource()...)
	}
	return tc
}
