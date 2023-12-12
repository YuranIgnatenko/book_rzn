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

func NewTableOrdersHistory(db *sql.DB) *TableOrders {
	return &TableOrders{DB: db}
}

func (t_orders_history *TableOrdersHistory) DeleteTableOrdersHistory(tokenUser, idOrder string) {
	rows, err := t_orders_history.DB.Query(fmt.Sprintf(`DELETE FROM bookrzn.OrdersHistory WHERE token="%s" AND id_order="%s";`, tokenUser, idOrder))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

// сделать вставку строки в OrdersHistory
func (t_orders_history *TableOrdersHistory) SaveTargetInOrdersHistory(db *sql.DB, token, id_order string) {
	fmt.Println("start SAVE orders history________-->:", token, id_order, "token, id order")

	var t_hash, t_count string
	var mapa_hash_order = make(map[string]string, 0)

	rows, err := t_orders_history.DB.Query(`SELECT target_hash,count FROM bookrzn.Orders WHERE token='` + token + `' AND id_order='` + id_order + `';`)
	if err != nil {
		panic(err)
	}
	fmt.Println("check ", err, rows, len(mapa_hash_order), mapa_hash_order)
	defer rows.Close()
	for rows.Next() {
		fmt.Scan(&t_hash, &t_count)
		mapa_hash_order[t_hash] = t_count
		fmt.Println("selecting ----->>>", t_hash, t_count, len(mapa_hash_order), mapa_hash_order)
		if err != nil {
			panic(err)
		}
	}

	for target_hash, target_count := range mapa_hash_order {
		fmt.Println("INSERTING --->>>", target_hash, target_count, id_order)
		rows, err = t_orders_history.DB.Query(
			fmt.Sprintf(`INSERT INTO bookrzn.OrdersHistory (token,target_hash,count,date,id_order,status_order)
		VALUES ( '%s','%s','%s','%s','%s','%s');`,
				token, target_hash, target_count, DateNow(), id_order, "on"))
		if err != nil {
			panic(err)
		}
		defer rows.Close()

	}
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

		fc, err := strconv.ParseFloat(card.Count, 64)
		if err != nil {
			fc = 0.1
		}
		temp_fp := strings.ReplaceAll(card.Price, "\u00a0", "")
		fp, err := strconv.ParseFloat(temp_fp, 64)
		if err != nil {
			// panic(err)
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
			panic(err)
		}
	}
	return card
}
