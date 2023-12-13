package connector

import (
	"backend/models"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// import (
// 	"backend/tools"
// 	"fmt"
// 	"strings"
// 	"time"
// )

// // Naming functions

func DateNow() string {
	t := time.Now()
	ts := strings.ReplaceAll(strings.Split(fmt.Sprintf("%v", t), " ")[0], "-", ".")
	return ts
}

func DateTimeNow() string {
	t := time.Now()
	// ts := strings.ReplaceAll(strings.Split(fmt.Sprintf("%v", t), " ")[0], "-", ".")
	return fmt.Sprintf("%v", t)
}

// FROM bookrzn.Targets TO models.TargetCard (need targetHash)
func TargetCardFromTargetHash(db *sql.DB, target_hash string) models.TargetCard {
	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Targets WHERE target_hash = '%s';`, target_hash)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	card := models.TargetCard{}

	for rows.Next() {
		card := models.TargetCard{}
		err := rows.Scan(
			&card.Id,
			&card.TargetHash,
			&card.Autor,
			&card.Title,
			&card.Price,
			&card.Link,
			&card.Comment,
			&card.Source,
			&card.Tag,
		)

		if err != nil {
			panic(err)
		}
	}
	return card
}

// //return 'target_hash' from "Orders" --> []string
// func GetTargetHashStringList(token string) []string {
// 	list := make([]string, 0)
// 	temp := ""

// 	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT target_hash FROM bookrzn.Orders WHERE token='%s';`, token)) //,
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		rows.Scan(&temp)

// 		if err != nil {
// 			continue
// 		}
// 		list = append(list, temp)
// 	}

// 	return list
// }

// // retrun []string from 'keys' to map['key']..
// func ConvMapToString(data map[string]int) []string {
// 	data_temp := make([]string, 0)
// 	for k, _ := range data {
// 		data_temp = append(data_temp, k)
// 	}
// 	return data_temp
// }

// //return 'column' from "table" --> []string
// func GetListFromColumnTable(column, table string) []string {
// 	data_temp := make(map[string]int, 0)

// 	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT %s FROM bookrzn.%s;`, column, table))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	var temp string

// 	for rows.Next() {
// 		if err != nil {
// 			continue
// 		}
// 		rows.Scan(&temp)
// 		data_temp[temp] = 0
// 	}
// 	data_list := conn.ConvMapToString(data_temp)

// 	return data_list
// }

// // //return 'column' from "table" --> []string
// // func  GetListMapaFromColumnTable(column, table string) []string {
// // 	data_temp := make(map[string]int, 0)

// // 	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT %s FROM bookrzn.%s;`, column, table))
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	defer rows.Close()

// // 	var temp string

// // 	for rows.Next() {
// // 		if err != nil {
// // 			continue
// // 		}
// // 		rows.Scan(&temp)
// // 		data_temp[temp] = 0
// // 	}
// // 	data_list := conn.ConvMapToString(data_temp)

// // 	return data_list
// // }

// // //return 'all' from "table" --> map[string]string
// // func  GetMapaFromColumnTable(table string) map[string]string {
// // 	data_temp := make(map[string]int, 0)
// // 	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT * FROM bookrzn.%s;`, table))
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	defer rows.Close()
// // 	var temp string
// // 	for rows.Next() {
// // 		if err != nil {
// // 			continue
// // 		}
// // 		rows.Scan(&temp)
// // 		data_temp[temp] = 0
// // 	}
// // 	data_list := conn.ConvMapToString(data_temp)
// // 	return data_list
// // }

// // //return 'token' from "Orders" --> []string
// // func  ListTokenUsersFromOrders() []string {
// // 	return conn.GetListFromColumnTable("token", "Orders")
// // }

// // retrun random token "234567543567543" --> string
// func GetNewRandomNumberFastOrder() string {
// 	return tools.NewRandomTokenFastOrder()
// }
