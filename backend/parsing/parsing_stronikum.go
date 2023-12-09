package parsing

import (
	"backend/models"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type ServiceStronikum struct {
	LinksVisit  []string
	TargetCards []models.TargetCard
	SourceType  string
	TagName     string
}

func NewServiceStronikum(links []string, tagname string) *ServiceStronikum {
	return &ServiceStronikum{
		LinksVisit: links,
		TagName:    tagname,
		SourceType: "stronikum.ru",
	}
}

func (ss *ServiceStronikum) ScrapSource() []models.TargetCard {
	mp_title, mp_price, mapa_links := ss.parseTitlesPrices()
	mp_image := ss.parseImages(mapa_links, mp_title)
	dt := ss.targetCardsFromMapa(mp_title, mp_price, mp_image)
	return dt
}

func (ss *ServiceStronikum) targetCardsFromMapa(mp_title, mp_price, mp_image map[int]string) []models.TargetCard {
	array_target_cards := make([]models.TargetCard, 0)
	for ind, _ := range mp_title {
		tc := models.TargetCard{
			Autor:   "-",
			Title:   mp_title[ind],
			Link:    mp_image[ind],
			Price:   mp_price[ind],
			Tag:     ss.TagName,
			Comment: "desc",
			Source:  ss.SourceType,
		}
		fmt.Println("Tag:    ", ss.TagName)

		tc.Id = "-" + mp_title[ind] + mp_price[ind] + "https://stronikum.ru" + mp_image[ind]
		tc.TargetHash = fmt.Sprintf("%v", time.Now().UnixNano())

		array_target_cards = append(array_target_cards, tc)
	}
	return array_target_cards
}

// retrun data-titles, data-prices data-links
func (ss *ServiceStronikum) parseTitlesPrices() (map[int]string, map[int]string, map[int]string) {
	c := colly.NewCollector()
	mapa_id_to_title := make(map[int]string, 0)
	mapa_id_to_price := make(map[int]string, 0)
	mapa_id_to_link := make(map[int]string, 0)
	flagSwitchTitlePrice := 1 // 1-title\ 2-price
	id_temp_title := 0
	id_temp_price := 0
	id_temp_link := 0

	for _, link := range ss.LinksVisit {
		c.OnHTML(".type-price", func(e *colly.HTMLElement) {
			e.ForEach("td", func(_ int, el *colly.HTMLElement) {
				flagRun := true
				if string(el.Text[0]) == "\n" {
					flagRun = false
				}

				if flagRun {
					text := strings.TrimSpace(el.Text)

					if flagSwitchTitlePrice == 1 {
						mapa_id_to_title[id_temp_title] = text
						id_temp_title += 1
						flagSwitchTitlePrice = 2
					} else if flagSwitchTitlePrice == 2 {
						text = strings.ReplaceAll(text, " р.", "")
						mapa_id_to_price[id_temp_price] = text
						id_temp_price += 1
						flagSwitchTitlePrice = 1
					}
					mapa_id_to_link[id_temp_link] = link
					id_temp_link += 1
				}

			})
		})
		c.Visit(link)
	}
	return mapa_id_to_title, mapa_id_to_price, mapa_id_to_link
}

// retrun data-titles, data-prices
func (ss *ServiceStronikum) parseImages(data_links, data_title map[int]string) map[int]string {
	mapa_id_to_image := make(map[int]string, 0)

	c := colly.NewCollector()
	index := 0
	for _, link := range data_links {
		c.OnHTML("a", func(e *colly.HTMLElement) {

			url := strings.TrimSpace(e.Attr("href"))
			flagRun := true
			if string(url[0]) == "#" || string(url[0]) == "/" || url == "mailto:info@stronikum.ru" {
				flagRun = false
			}
			if flagRun {
				code := strings.Split(strings.Split(url, "/")[1], "_")[0]
				s := "https://stronikum.ru/photo/" + code + ".jpg"
				mapa_id_to_image[index] = s
				index += 1
			}
		})
		c.Visit(link)
	}
	return mapa_id_to_image
}
