package parsing

import (
	"backend/models"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type ItemNaura struct {
	NameKabinet string
	NamePredmet string
	NameTitle   string
	Link        string
}
type ServiceNaura struct {
	LinksVisit  []string
	TargetCards []models.TargetCard
	SourceType  string
	TagName     string
}

func NewServiceNaura(links []string, tagname string) *ServiceNaura {
	return &ServiceNaura{
		LinksVisit: links,
		TagName:    tagname,
		SourceType: "nau-ra.ru",
	}
}

func (sn *ServiceNaura) ScrapSource() []models.TargetCard {
	data_mapa_parts_links := sn.getPartsLinks()
	data_mapa_parts_images := sn.getLinksImages(data_mapa_parts_links)
	data_mapa_parts_desc := sn.getPartsDesc(data_mapa_parts_links)

	targets := sn.convertToTargetCards(data_mapa_parts_links, data_mapa_parts_images, data_mapa_parts_desc)
	return targets
}

func (sn *ServiceNaura) getPartsDesc(dt map[string]string) map[string][]string {
	mapa_parts_desc := make(map[string][]string, 0)
	c := colly.NewCollector()
	for part, link := range dt {
		// подразделы 804
		c.OnHTML(".categories", func(e *colly.HTMLElement) {
			e.ForEach("p", func(_ int, el *colly.HTMLElement) {
				mapa_parts_desc[part] = append(mapa_parts_desc[part], el.Text)
			})
		})

		c.Visit(link)
	}

	return mapa_parts_desc
}

func (sn *ServiceNaura) getLinksImages(dt map[string]string) map[string][]string {
	mapa_parts_images := make(map[string][]string, 0)
	c := colly.NewCollector()
	for parts, link := range dt {
		// подразделы 804
		flagRun := true
		c.OnHTML(".img-wrapper", func(e *colly.HTMLElement) {
			e.ForEach("img[src]", func(_ int, el *colly.HTMLElement) {
				if strings.TrimSpace(string(el.Attr("src"))) != "" {
					for _, elem := range mapa_parts_images[parts] {
						if elem == "https://nau-ra.ru/"+el.Attr("src") {
							flagRun = false
						}
					}

					if flagRun {
						mapa_parts_images[parts] = append(mapa_parts_images[parts], "https://nau-ra.ru/"+el.Attr("src"))
					}
				}
			})
		})
		c.Visit(link)
	}
	return mapa_parts_images
}

func (sn *ServiceNaura) getPartsLinks() map[string]string {
	mapa_parts_links := make(map[string]string, 0)
	list_parts := make([]string, 0)
	list_links := make([]string, 0)

	c := colly.NewCollector()
	for _, link := range sn.LinksVisit {
		// разделы 804
		c.OnHTML(".blog-content", func(e *colly.HTMLElement) {
			e.ForEach("a[href]", func(_ int, el *colly.HTMLElement) {
				// fmt.Println(el.Text)
				list_parts = append(list_parts, el.Text)
			})
		})
		// ссылки на разделы 804
		c.OnHTML(".blog-content", func(e *colly.HTMLElement) {
			e.ForEach("a[href]", func(_ int, el *colly.HTMLElement) {
				list_links = append(list_links, el.Attr("href"))

			})
		})

		c.Visit(link)
	}

	for i, elem := range list_parts {
		if elem == "Приказом Министерства просвещения РФ № 804 от 06.09.2022 года" {
			continue
		}
		elem = strings.Join(strings.Split(elem, " ")[1:], " ")
		mapa_parts_links[elem] = list_links[i]

	}
	return mapa_parts_links
}

func (sn *ServiceNaura) convertToTargetCards(
	dt_parts_links map[string]string,
	dt_parts_images map[string][]string,
	dt_parts_desc map[string][]string) []models.TargetCard {

	list_target_cards := make([]models.TargetCard, 0)

	for part, img := range dt_parts_images {

		dt := models.TargetCard{
			Autor:   "-",
			Title:   part,
			Link:    img[0],
			Price:   "--",
			Comment: strings.Join(dt_parts_desc[part], "\n"),
			Tag:     sn.TagName,
			Source:  sn.SourceType,
		}

		dt.Title = strings.ReplaceAll(dt.Title, `"`, "")
		dt.Id = dt.Autor + dt.Title + dt.Price + "https://nau-ra.ru" + img[0]

		dt.TargetHash = fmt.Sprintf("%v", time.Now().UnixNano())

		list_target_cards = append(list_target_cards, dt)
	}
	return list_target_cards
}
