package connector

import (
	"backend/models"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type TableOrders struct {
	DB *sql.DB
}

func NewTableOrders() *TableOrders {
	return &TableOrders{}
}

// получить "count" из Orders-table from target_hash, id_order
func (t_orders *TableOrders) GetCountFromIdOrderHash(id_order string) string {
	rows, err := t_orders.DB.Query(fmt.Sprintf(`SELECT count FROM bookrzn.Orders WHERE id_order = '%s';`, id_order))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {

		var temp_count string

		rows.Scan(
			&temp_count,
		)
		if err != nil {
			panic(err)
		}
		return temp_count
	}
	return "error count"
}

func (t_orders *TableOrders) SearchTargetList(request string) []models.TargetCard {
	var target_search []models.TargetCard

	rows, err := t_orders.DB.Query(`SELECT * FROM bookrzn.Orders WHERE target_hash LIKE '%` + request + `%' ;`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var ts models.TargetCard
		rows.Scan(
			&ts.TargetHash,
			&ts.Autor,
			&ts.Price,
			&ts.Title,
			&ts.Price,
			&ts.Link,
			&ts.Comment,
		)

		if err != nil {
			panic(err)
		}
		target_search = append(target_search, ts)
	}
	return target_search
}

// FROM token IN BD Orders TO map[token][]string{[]row1,[]row2,[]row3,..}
func (t_orders *TableOrders) DataFromOrders(token string) map[string][][]string {
	data_map := make(map[string][][]string, 0)
	rows, err := t_orders.DB.Query(fmt.Sprintf(
		`SELECT (target_hash,count,date,id_order) 
		FROM bookrzn.Orders WHERE token LIKE %s ;`, token))

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var target_hash, count, date, id_order string

		rows.Scan(
			&target_hash,
			&count,
			&date,
			&id_order,
		)
		if err != nil {
			panic(err)
		}
		temp_str := []string{target_hash, count, date, id_order}
		data_map[token] = append(data_map[token], temp_str)
	}

	return data_map
}

func (t_orders *TableOrders) DeleteTableOrders(tokenUser, idOrder string) {
	rows, err := t_orders.DB.Query(fmt.Sprintf(`DELETE FROM bookrzn.Orders WHERE token="%s" AND id_order="%s";`, tokenUser, idOrder))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

func (t_orders *TableOrders) TargetCardsFromListOrders(token string) []models.TargetCard {
	mapa_target_hash_count := make(map[string]string, 0)
	mapa_target_hash_date := make(map[string]string, 0)
	mapa_target_hash_id_order := make(map[string]string, 0)
	list_target_hash := make([]string, 0)
	temp_target_cards_all := make([]models.TargetCard, 0)
	main_target_cards_all := make([]models.TargetCard, 0)
	temp_token, temp_target_hash, temp_count, temp_date, temp_id_order := "", "", "", "", ""

	rows, err := t_orders.DB.Query(fmt.Sprintf(`SELECT token,target_hash,count,date,id_order FROM bookrzn.Orders WHERE token='%s';`, token)) //,
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
		temp_target_cards_all = append(temp_target_cards_all, TargetCardFromTargetHash(t_orders.DB, hash))
	}

	for _, card := range temp_target_cards_all {
		card.Count = mapa_target_hash_count[card.TargetHash]
		card.Date = mapa_target_hash_date[card.TargetHash]
		card.IdOrder = mapa_target_hash_id_order[card.TargetHash]
		card.Price = strings.ReplaceAll(card.Price, ",", ".")
		fc, err := strconv.ParseFloat(card.Count, 64)
		if err != nil {
			panic(err)
		}
		temp_fp := strings.ReplaceAll(card.Price, "\u00a0", "")
		fp, err := strconv.ParseFloat(temp_fp, 64)
		if err != nil {
			panic(err)
		}
		card.Summa = float64(fc*fp) + 0.0
		main_target_cards_all = append(main_target_cards_all, card)
	}

	return main_target_cards_all
}
