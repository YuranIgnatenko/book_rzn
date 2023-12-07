package connector

import (
	"backend/models"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type TableOrdersHistory struct {
	DB *sql.DB
}

func NewTableOrdersHistory() *TableOrders {
	return &TableOrders{}
}

// сделать вставку строки в OrdersHistory
func (t_orders_history *TableOrdersHistory) SaveTargetInOrdersHistory(token, target_hash, count string, id_order string) {
	var rows *sql.Rows
	var err error

	rows, err = t_orders_history.DB.Query(
		fmt.Sprintf(`INSERT INTO bookrzn.OrdersHistory (token, target_hash, count, date, id_order, status_order) 
		VALUES ( '%s','%s','%s','%s','%s','%s');`,
			token, target_hash, count, DateNow(), id_order, "on"))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

// получить массив карточек товара по токену юзера
func (t_orders_history *TableOrdersHistory) TargetCardsFromListOrdersHistory(token string) []models.TargetCard {
	var (
		mapa_target_hash_count    = make(map[string]string, 0)
		mapa_target_hash_date     = make(map[string]string, 0)
		mapa_target_hash_id_order = make(map[string]string, 0)
		list_target_hash          = make([]string, 0)
		temp_target_cards_all     = make([]models.TargetCard, 0)
		main_target_cards_all     = make([]models.TargetCard, 0)
	)
	temp_token, temp_target_hash, temp_count, temp_date, temp_id_order := "", "", "", "", ""

	rows, err := t_orders_history.DB.Query(fmt.Sprintf(
		`SELECT token,target_hash,count,date,id_order,status_order 
		FROM bookrzn.OrdersHistory WHERE token='%s';`, token))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&temp_token, &temp_target_hash, &temp_count, &temp_date, &temp_id_order)

		if err != nil {
			panic(err)
		}
		list_target_hash = append(list_target_hash, temp_target_hash)
		mapa_target_hash_count[temp_target_hash] = temp_count
		mapa_target_hash_date[temp_target_hash] = temp_date
		mapa_target_hash_id_order[temp_target_hash] = temp_id_order
	}

	for _, hash := range list_target_hash {
		temp_target_cards_all = append(temp_target_cards_all, t_orders_history.TargetCardFromTargetHash(hash))
	}

	for _, card := range temp_target_cards_all {
		card.Count = mapa_target_hash_count[card.TargetHash]
		card.Date = mapa_target_hash_date[card.TargetHash]
		card.IdOrder = mapa_target_hash_id_order[card.TargetHash]
		card.Price = strings.ReplaceAll(card.Price, ",", ".")

		fmt.Println()
		fmt.Println("card.Count,card.Date,card.IdOrder,card.Price", card.Count, card.Date, card.IdOrder, card.Price)
		fmt.Println()

		fc, err := strconv.ParseFloat(card.Count, 64)
		if err != nil {
			fc = 0.1
			fmt.Println("card.Count=", card.Count)
			// panic(err)
		}
		temp_fp := strings.ReplaceAll(card.Price, "\u00a0", "")
		fp, err := strconv.ParseFloat(temp_fp, 64)
		if err != nil {
			// panic(err)
			fmt.Println("temp_fp=", temp_fp)
			fp = 0.2
		}
		card.Summa = float64(fc * fp)
		main_target_cards_all = append(main_target_cards_all, card)
	}

	return main_target_cards_all
}

func (t_orders_history *TableOrdersHistory) TargetCardFromTargetHash(target_hash string) models.TargetCard {
	rows, err := t_orders_history.DB.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Targets WHERE target_hash = '%s';`, target_hash)) //,
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
			continue
		}
		return card
	}
	return card
}
