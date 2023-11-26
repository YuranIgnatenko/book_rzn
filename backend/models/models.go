package models

type ServiceScraper interface {
	ScrapSource() []TargetCard
}

type OrderCMS struct {
	Name     string
	Date     string
	Phone    string
	Email    string
	CountAll string
	PriceAll string
}

type TargetCard struct {
	Autor      string //`csv:"autor"`
	Title      string //`csv:"title"`
	Price      string //`csv:"price"`
	Link       string //`csv:"link"`
	Id         string
	Comment    string
	TargetHash string
	Source     string
	Tag        string
}

type OrdersRows struct {
	Link         string //`csv:"link"`
	Id           string
	Comment      string
	Price        string //`csv:"price"`
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
