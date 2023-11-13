package tools

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsvFile(filePath string) [][]string {
	s := make([][]string, 0)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
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
