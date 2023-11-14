package tools

// func ReadCsvFile(filePath string) [][]string {
// 	s := make([][]string, 0)

// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	reader.FieldsPerRecord = 3
// 	reader.Comment = '#'

// 	for {
// 		record, e := reader.Read()
// 		if e != nil {
// 			break
// 		}
// 		s = append(s, record)
// 	}
// 	return s

// }
