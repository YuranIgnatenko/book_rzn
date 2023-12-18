package models

import (
	"fmt"
)

type PageTarget struct {
	PageNext     int
	PagePrev     int
	PageOld      int
	PageTotal    int
	PageData     []TargetCard
	PageDataAll  []TargetCard
	PageSize     int //recmmended 20 cards targets
	PageNow      int
	PageLinkNext string
	PageLinkPrev string
	// LastSearch   string
	// Page
}

// TODO: pass func block
func (pt *PageTarget) SortedBySwitch(filter string) []TargetCard {
	return []TargetCard{}
}

func (pt *PageTarget) GetPage(link string, number_page int) []TargetCard {
	// math.Ceil() — округление в большую сторону
	// math.Floor() - в меньшую
	// page_total := math.Ceil(pt.PageTotal / pt.PageSize)

	count_pages := int(pt.PageTotal/pt.PageSize) + 1
	pages := make(map[int][]TargetCard, count_pages)
	// fmt.Println("countPage :", pages)

	temp_counter_size := 0
	temp_ind := 1
	// all_count := 0

	for _, card := range pt.PageDataAll {
		// all_count = ind
		if temp_counter_size == pt.PageSize {
			temp_counter_size = 0
			temp_ind += 1
		}

		pages[temp_ind] = append(pages[temp_ind], card)
		temp_counter_size += 1
	}

	// fmt.Println("list_pages max page:", ))
	pt.PageNow = number_page
	pt.PagePrev = number_page - 1
	pt.PageNext = number_page + 1
	pt.PageData = pages[number_page]
	pt.PageTotal = len(pages)

	if pt.PageNext > pt.PageTotal {
		pt.PageNext = pt.PageTotal
	}
	if pt.PagePrev < 1 {
		pt.PagePrev = 1
	}
	pt.PageLinkNext = "/" + link + "/" + fmt.Sprint(pt.PageNext)
	pt.PageLinkPrev = "/" + link + "/" + fmt.Sprint(pt.PagePrev)

	fmt.Println("pt.PageLinkNext,pt.PageLinkPrev::", pt.PageLinkNext, pt.PageLinkPrev)

	return pages[number_page]

}

func (pt *PageTarget) GetPageSearch(link string, number_page int) []TargetCard {
	// math.Ceil() — округление в большую сторону
	// math.Floor() - в меньшую
	// page_total := math.Ceil(pt.PageTotal / pt.PageSize)

	count_pages := int(pt.PageTotal/pt.PageSize) + 1
	pages := make(map[int][]TargetCard, count_pages)
	// fmt.Println("countPage :", pages)

	temp_counter_size := 0
	temp_ind := 1
	// all_count := 0

	for _, card := range pt.PageDataAll {
		// all_count = ind
		if temp_counter_size == pt.PageSize {
			temp_counter_size = 0
			temp_ind += 1
		}

		pages[temp_ind] = append(pages[temp_ind], card)
		temp_counter_size += 1
	}

	// fmt.Println("list_pages max page:", ))
	pt.PageNow = number_page
	pt.PagePrev = number_page - 1
	pt.PageNext = number_page + 1
	pt.PageData = pages[number_page]
	pt.PageTotal = len(pages)

	if pt.PageNext > pt.PageTotal {
		pt.PageNext = pt.PageTotal
	}
	if pt.PagePrev < 1 {
		pt.PagePrev = 1
	}
	pt.PageLinkNext = "/" + link + "/" + fmt.Sprint(pt.PageNext)
	pt.PageLinkPrev = "/" + link + "/" + fmt.Sprint(pt.PagePrev)

	return pages[number_page]

}

type ServiceScraper interface {
	ScrapSource() []TargetCard
}

// type OrdersCardsCms struct {
// 	Name     string
// 	Date     string
// 	Phone    string
// 	Email    string
// 	Token    string
// 	Targets  []TargetCard
// 	CountAll string
// 	PriceAll string
// }

type ListOrdersTargetCard struct {
	Orders              map[string][]TargetCard // {"1":[]TargetCard{...}, "2":[]TargetCard{...}} --> key=id_order from bookrzn.Orders
	OrdersStatusConfirm map[string]bool
	PriceFinish         map[string]float64
}

type MenuCard struct {
	Title  string
	Link   string
	PathTo string
}

type TargetCard struct {
	Autor                string
	Title                string
	Price                string
	Link                 string
	Id                   string
	Comment              string
	TargetHash           string
	Source               string // site.com
	Tag                  string
	Count                string
	Summa                float64 //f.ff - 0.00
	Date                 string  // YYYY.mm.dd
	IdOrder              string
	Status               string
	Edition_order_token  string
	CMSNameOrders        string
	CMSDateOrders        string
	CMSPhoneOrders       string
	CMSEmailOrders       string
	CMSTokenOrders       string
	CMSTargetsHashOrders []string
	CMSCountAllOrders    string
	CMSPriceAllOrders    string
}

type OrdersRows struct {
	Link         string
	Id           string
	Comment      string
	Price        string
	CountTargets string
}

type OrdersCards struct {
	Autor      string //`csv:"autor"`
	Title      string //`csv:"title"`
	Price      string //`csv:"price"`
	Link       string //`csv:"link"`
	Id         string
	Comment    string
	TargetHash string
}

type FavoritesCards struct {
	Autor      string //`csv:"autor"`
	Title      string //`csv:"title"`
	Price      string //`csv:"price"`
	Link       string //`csv:"link"`
	Id         string
	Comment    string
	TargetHash string
}

type Users struct {
	Id              int
	Login, Password string
	Token, Name     string
	Phone, Family   string
	Type, Email     string
}

type Favorites struct {
	Autor      string //`csv:"autor"`
	Title      string //`csv:"title"`
	Price      string //`csv:"price"`
	Link       string //`csv:"link"`
	Id         string
	Comment    string
	TargetHash string
	Count      string
	Source     string // site.com
	Tag        string
}

type Orders struct {
	Id         int
	Token      string
	TargetHash string
	Count      string
	Datetime   string
}

type Targets struct {
	Id         int
	TargetHash string
	Autor      string
	Title      string
	Price      string
	Image      string
	Comment    string
}

type DataFastOrder struct {
	Name            string   `json:"Name"`
	Phone           string   `json:"Phone"`
	Email           string   `json:"Email"`
	ArrTarget       []string `json:"ArrTarget"`
	ArrTargetCount  []string `json:"ArrTargetCount"`
	NumberFastOrder string   `json:"NumberFastOrder"`
}
type DataFastOrderOne struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Target string `json:"target"`
	Count  string `json:"count"`
	Token  string `json:"token"`
}
