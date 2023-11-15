package core

type ProsvCarder interface{}

type ParsingServicer interface {
	ScrapSource() []ProsvCarder
	WriteToCsv(data []*ProsvCarder)
}
