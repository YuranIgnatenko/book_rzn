package models

type TargetCarder interface{}

type ParsingServicer interface {
	ScrapSource() []TargetCarder
	WriteToCsv(data []*TargetCarder)
}

type TargetCard struct {
	Autor      string //`csv:"autor"`
	Title      string //`csv:"title"`
	Price      string //`csv:"price"`
	Link       string //`csv:"link"`
	Id         string
	Comment    string
	TargetHash string
}

// type ProsvCard struct {
// 	Autor string //`csv:"autor"`
// 	Title string //`csv:"title"`
// 	Price string //`csv:"price"`
// 	Link  string //`csv:"link"`
// 	Id    string
// }

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
