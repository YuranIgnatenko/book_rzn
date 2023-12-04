package models

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
