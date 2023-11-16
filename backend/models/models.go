package models

type ProsvCarder interface{}

type ParsingServicer interface {
	ScrapSource() []ProsvCarder
	WriteToCsv(data []*ProsvCarder)
}

type ProsvCard struct {
	Autor string //`csv:"autor"`
	Title string //`csv:"title"`
	Price string //`csv:"price"`
	Link  string //`csv:"link"`
	Id    string
}

type FavoritesCards struct {
	Autor string //`csv:"autor"`
	Title string //`csv:"title"`
	Price string //`csv:"price"`
	Link  string //`csv:"link"`
	Id    string
}

type Users struct {
	Id              int
	Login, Password string
	Token, Name     string
	Phone, Family   string
	Type, Email     string
}

type Favorites struct {
	Id         int
	Token      string
	TargetHash string
	Count      string
	Datetime   string
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
