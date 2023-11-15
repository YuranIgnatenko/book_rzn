package bd

import (
	"backend/config"
	"encoding/csv"
	"os"
)

type Bd struct {
	config.Configuration
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

func (b *Bd) SaveParsingProsv() {}

func (b *Bd) ReadConfigPLatform() {}
