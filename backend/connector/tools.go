package connector

import (
	"backend/models"
	"backend/tools"
	"fmt"
)

// Naming functions

// get 'models.TargetCard' from 'target_hash' bd 'bookrzn.Targets'
func (conn *Connector) TargetCardFromTargetHash(target_hash string) models.TargetCard {
	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Targets WHERE target_hash = '%s';`, target_hash)) //,
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
			fmt.Println(err)
			continue
		}
		return card
	}
	return card
}

//return map["tokenUser"][]string{"targetHash","countTarget"}
func (conn *Connector) MapaTokenUserToOrders(token_user string) map[string][]string {
	var (
		data                                     = make(map[string][]string, 0)
		temp_token, temp_target_hash, temp_count string
	)

	rows, err := conn.Db.Query(
		fmt.Sprintf(`SELECT token,target_hash,count FROM bookrzn.Orders WHERE token='%s';`, token_user)) //,
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&temp_token, &temp_target_hash, &temp_count)

		if err != nil {
			fmt.Println(err)
			continue
		}
		data[token_user] = append(data[token_user], temp_target_hash)
	}
	return data
}

//return 'target_hash' from "Orders" --> []string
func (conn *Connector) GetTargetHashStringList(token string) []string {
	list := make([]string, 0)
	temp := ""

	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT target_hash FROM bookrzn.Orders WHERE token='%s';`, token)) //,
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&temp)

		if err != nil {
			fmt.Println(err)
			continue
		}
		list = append(list, temp)
	}

	return list
}

//return 'token' from "Orders" --> []string
func (conn *Connector) ListTokenUsersFromOrders() []string {
	data_list_users_tokens := make([]string, 0)
	isNewToken := func(token1 string) bool {
		for _, token2 := range data_list_users_tokens {
			if token2 == token1 {
				return false
			}
		}
		return true
	}

	rows, err := conn.Db.Query(`SELECT token FROM bookrzn.Orders;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var temp_token string

	for rows.Next() {
		rows.Scan(&temp_token)
		isRes := isNewToken(temp_token)
		if !isRes {
			continue
		}

		if err != nil {
			fmt.Println(err)
			continue
		}
		data_list_users_tokens = append(data_list_users_tokens, temp_token)
	}
	return data_list_users_tokens
}

// retrun random token "234567543567543" --> string
func (conn *Connector) GetNewRandomNumberFastOrder() string {
	return tools.NewRandomTokenFastOrder()
}

// func (conn *Connector) GetListOrdersRow(token string) []models.OrdersRows {
// 	db, err := sql.Open("mysql", conn.dsn())
// 	if err != nil {
// 		fmt.Printf("Error %s when opening DB\n", err)
// 	}
// 	conn.Db = db
// 	list_targets_hash := []string{}
// 	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT target_hash FROM bookrzn.Orders WHERE token='%s';`, token))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	var ordcards []models.OrdersRows
// 	for rows.Next() {
// 		var target_hash string
// 		rows.Scan(&target_hash)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		list_targets_hash = append(list_targets_hash, target_hash)
// 	}
// 	for _, hash := range list_targets_hash {
// 		rows, err = conn.Db.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Targets WHERE target_hash='%s';`, hash))
// 		if err != nil {
// 			panic(err)
// 		}
// 		for rows.Next() {
// 			card := models.OrdersRows{}
// 			err := rows.Scan(
// 				&card.Id,
// 				&card.Price,
// 				&card.Link,
// 				&card.Comment,
//
// 			if err != nil {
// 				fmt.Println(err)
// 				continue
// 			}
// 			ordcards = append(ordcards, card)
// 		}
// 	}
// 	return ordcards
// }

// func (conn *Connector) FastCardsFromFastOrders() []models.DataFastOrderOne {
// 	var rows *sql.Rows
// 	var err error

// 	dt := make([]models.DataFastOrderOne, 0)

// 	rows, err = conn.Db.Query(`SELECT * FROM bookrzn.FastOrders ;`)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for rows.Next() {
// 		dtfast := models.DataFastOrderOne{}
// 		temp := ""

// 		err := rows.Scan(
// 			&temp,
// 			&dtfast.Name,
// 			&dtfast.Phone,
// 			&dtfast.Email,
// 			&dtfast.Target,
// 			&dtfast.Count,
// 			&dtfast.Token,
// 		)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		dt = append(dt, dtfast)
// 	}

// 	// обязательно иначе привысит лимит подключений и будет сбой
// 	defer rows.Close()
// 	return dt
// }

// //return new 'id' for table "Num" --> string
// func (conn *Connector) GetNewNumberFastOrder() string {
// 	rows, err := conn.Db.Query(`SELECT MAX(Id) FROM bookrzn.FastOrders;`)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var value float64
// 	rows.Scan(&value)

// 	value += 1
// 	defer rows.Close()
// 	return fmt.Sprintf("%v", value)
// }

// func (conn *Connector) CountRows(namebd string) int {
// 	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT COUNT(*) FROM %s;`, namebd)) //,
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	var count string

// 	for rows.Next() {
// 		err := rows.Scan(&count)

// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}

// 	}
// 	c, err := strconv.Atoi(count)
// 	if err != nil {
// 		panic(err)

// 	}
// 	return c
// }

// func (conn *Connector) GetUserTokenList(token string) []string {}
// func (conn *Connector) GetTargetList(token string) []string {}
// func (conn *Connector) GetTargetList(token string) []string {}
// func (conn *Connector) GetTargetList(token string) []string {}
// func (conn *Connector) GetTargetList(token string) []string {}
// func (conn *Connector) GetTargetList(token string) []string {}
// func (conn *Connector) GetTargetList(token string) []string {}
