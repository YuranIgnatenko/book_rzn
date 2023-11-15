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
