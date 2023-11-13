package bd

import (
	"backend/config"
	"encoding/csv"
	"fmt"
	"os"
)

var Config = config.NewConfiguration()

func readCsvFileRows(filePath string) [][]string {
	s := make([][]string, 0)

	file, err := os.Open(filePath)

	fmt.Println("||||||||||", err, filePath)
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
			// fmt.Println(e)
			break
		}
		s = append(s, record)
		fmt.Println(record)
	}
	return s

}

func ReadUsersData() [][]string {
	rows := readCsvFileRows(Config.Path_bd + Config.Bd_users_list)
	return rows
}

func ReadAdminData() [][]string {
	rows := readCsvFileRows(Config.Path_bd + Config.Bd_admin_list)
	return rows
}

func SaveParsingProsv() {

}

func ReadConfigPLatform() {}

func ReadOrdersList() {}
