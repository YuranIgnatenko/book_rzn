package bd

import (
	"backend/config"
	"backend/models"
	"encoding/csv"
	"fmt"
	"os"
)

type Bd struct {
	config.Configuration
}

func NewBd(c config.Configuration) *Bd {
	return &Bd{
		Configuration: c,
	}
}

func (b *Bd) readCsvFileRows(filePath string) [][]string {
	s := make([][]string, 0)

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 0
	reader.Comment = '#'

	for {
		record, e := reader.Read()
		if e != nil {
			break
		}
		s = append(s, record)
	}
	return s

}

func (b *Bd) ReadUsersData() [][]string {
	rows := b.readCsvFileRows(b.Path_bd + b.Bd_users_list)
	return rows
}

func (b *Bd) ReadAdminData() [][]string {
	rows := b.readCsvFileRows(b.Path_bd + b.Bd_admin_list)
	return rows
}

func (b *Bd) ReadFavorites() [][]string {
	rows := b.readCsvFileRows(b.Path_bd + b.Bd_favorites)
	return rows
}

func (b *Bd) ReadOrders() [][]string {
	rows := b.readCsvFileRows(b.Path_bd + b.Bd_orders)
	return rows
}

func (b *Bd) SaveTarget(token, id_target string) {
	fmt.Println("save target (in /add favorites)")

	file, err := os.OpenFile(b.Path_bd+b.Bd_favorites, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	row := []string{token, id_target}
	err = writer.Write(row)
	if err != nil {
		panic(err)
	}

}

func (b *Bd) FindTarget(token string) []models.ProsvCard {
	fmt.Println("find target (in /favorites)")
	data_tokens := make([]string, 0)

	file, err := os.Open(b.Path_bd + b.Bd_favorites)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 0
	reader.Comment = '#'
	rec_all, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, row := range rec_all {
		if row[0] == token {
			data_tokens = append(data_tokens, row[1])
		}
	}

	cards := make([]models.ProsvCard, 0)

	file, err = os.Open(b.Path_bd + b.Bd_prosv)
	if err != nil {
		// return nil, err
	}
	defer file.Close()
	reader = csv.NewReader(file)
	reader.FieldsPerRecord = 0
	reader.Comment = '#'
	rec_all, err = reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, row := range rec_all {
		if row[4] == token {
			cards = append(cards, models.ProsvCard{
				Autor: row[0],
				Title: row[1],
				Price: row[2],
				Link:  row[3],
				Id:    row[4],
			})

		}
	}
	fmt.Println(len(cards), "len cards")
	return cards
}

func (b *Bd) SaveParsingProsv() {}

func (b *Bd) ReadConfigPLatform() {}
