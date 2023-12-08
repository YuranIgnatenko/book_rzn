package connector

import (
	"backend/config"
	"backend/models"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Connector struct {
	config.Configuration
	Db                 *sql.DB
	TableOrders        TableOrders
	TableFavorites     TableFavorites
	TableOrdersHistory TableOrdersHistory
	TableTargets       TableTargets
	TableUsers         TableUsers
}

func (conn *Connector) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", conn.DB_user, conn.DB_password, conn.DB_hostname, conn.DB_database)
}

func NewConnector(c config.Configuration) *Connector {
	conn := Connector{
		Configuration: c,
	}

	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db
	conn.TableOrders.DB = db
	conn.TableOrdersHistory.DB = db
	conn.TableTargets.DB = db
	conn.TableFavorites.DB = db
	conn.TableUsers.DB = db

	return &conn
}

// ==============================================================================
// ==============================================================================
// ==============================================================================
// ==============================================================================
// ==============================================================================
// ==============================================================================

func (conn *Connector) DeleteTableOrdersHistory(tokenUser, idOrder string) {
	rows, err := conn.Db.Query(fmt.Sprintf(`DELETE FROM bookrzn.OrdersHistory WHERE token="%s" AND id_order="%s";`, tokenUser, idOrder))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

func (conn *Connector) TargetCardsFromToken() {

}

// FROM bookrzn.Targets TO models.TargetCard (need targetHash)
func (conn *Connector) TargetCardsFromOrders(target_hash string) models.TargetCard {
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
			panic(err)
		}
		return card
	}
	return card
}

// получить массив заказов из карточек товаров
func (conn *Connector) ListOrdersFromTargetCards(tc []models.TargetCard) models.ListOrdersTargetCard {
	res := models.ListOrdersTargetCard{}
	mapa_list_orders := make(map[string][]models.TargetCard, 0)
	mapa_prices := make(map[string]float64, 0)
	for _, card := range tc {
		mapa_list_orders[card.IdOrder] = append(mapa_list_orders[card.IdOrder], card)
	}

	for id_order, list_cards := range mapa_list_orders {
		for _, card := range list_cards {
			mapa_prices[id_order] += card.Summa
		}
	}
	res.PriceFinish = mapa_prices
	res.Orders = mapa_list_orders
	return res
}

// func (conn *Connector) TargetCardsFromListOrdersCMS() []models.TargetCard {
// 	main_target_cards_all := make([]models.TargetCard, 0)
// 	for _, token := range conn.GetListFromColumnTable("token", "Orders") {
// 		mapa_target_hash_count := make(map[string]string, 0)
// 		list_target_hash := make([]string, 0)
// 		temp_target_cards_all := make([]models.TargetCard, 0)
// 		temp_token, temp_target_hash, temp_count := "", "", ""

// 		rows, err := conn.Db.Query(fmt.Sprintf(`SELECT token,target_hash,count FROM bookrzn.Orders WHERE token='%s';`, token)) //,
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer rows.Close()

// 		for rows.Next() {
// 			rows.Scan(&temp_token, &temp_target_hash, &temp_count)

// 			if err != nil {
// 				continue
// 			}
// 			list_target_hash = append(list_target_hash, temp_target_hash)
// 			mapa_target_hash_count[temp_target_hash] = temp_count
// 		}

// 		for _, hash := range list_target_hash {
// 			temp_target_cards_all = append(temp_target_cards_all, conn.TargetCardFromTargetHash(hash))
// 		}

// 		for _, card := range temp_target_cards_all {
// 			user, err := conn.TableUsers.FindUserFromToken(token)
// 			if err != nil {
// 				panic(err)
// 			}

// 			card.CMSNameOrders = token
// 			card.CMSPhoneOrders = user.Phone
// 			card.CMSEmailOrders = user.Email
// 			// card.CMSPriceAllOrders = user.

// 			card.Count = mapa_target_hash_count[card.TargetHash]
// 			card.Price = strings.ReplaceAll(card.Price, ",", ".")
// 			// card.Title =
// 			// fc, err := strconv.ParseFloat(card.Count, 64)
// 			// if err != nil {
// 			// 	panic(err)
// 			// }
// 			// fp, err := strconv.ParseFloat(card.Price, 64)
// 			// if err != nil {
// 			// 	panic(err)
// 			// }
// 			// card.Price // float64(fc * fp)
// 			main_target_cards_all = append(main_target_cards_all, card)
// 		}
// 	}
// 	return main_target_cards_all
// }

func (conn *Connector) TargetCardsFromListFavorites(token string) []models.TargetCard {

	list_targets_hash := []string{}
	list_targets_count := []string{}

	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT target_hash,count FROM bookrzn.Favorites WHERE token='%s';`, token)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var favcards []models.TargetCard

	for rows.Next() {
		var target_hash string
		var target_count string

		rows.Scan(&target_hash, &target_count)

		if err != nil {
			panic(err)
		}
		list_targets_hash = append(list_targets_hash, target_hash)
		list_targets_count = append(list_targets_count, target_count)

	}

	for ind, hash := range list_targets_hash {
		rows, err = conn.Db.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Targets WHERE target_hash='%s';`, hash))
		if err != nil {
			panic(err)
		}

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
				&card.Tag,
				&card.Source,
			)
			card.Count = list_targets_count[ind]

			if err != nil {
				panic(err)
			}
			favcards = append(favcards, card)
		}
	}
	return favcards
}

func (conn *Connector) SaveTargetInOrders(token, target_hash, count string, id_order string) {
	var rows *sql.Rows
	var err error

	rows, err = conn.Db.Query(
		fmt.Sprintf(`INSERT bookrzn.Orders (token, target_hash, count, date, id_order) 
		VALUES ( '%s','%s','%s', '%s', '%s');`,
			token, target_hash, count, DateNow(), id_order))
	if err != nil {
		panic(err)
	}
	// обязательно иначе привысит лимит подключений и будет сбой
	defer rows.Close()
}

func (conn *Connector) SaveTargetInFavorites(token, targethash, count string) {
	rows, err := conn.Db.Query(
		fmt.Sprintf(`INSERT bookrzn.Favorites (token,target_hash,count) 
		VALUES ( '%s','%s','%s');`,
			token, targethash, count)) //,
	if err != nil {
		panic(err)
	}
	//
	// обязательно иначе привысит лимит подключений и будет сбой
	defer rows.Close()
}

func (conn *Connector) DeleteTargetFromFavorites(token, targethash string) error {

	rows, err := conn.Db.Query(
		fmt.Sprintf(`DELETE FROM bookrzn.Favorites
		WHERE token='%s' AND target_hash='%s';`, token, targethash))
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	return err
}

func (conn *Connector) AccessLogin(login, password string) string {
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db

	rows, err := conn.Db.Query(fmt.Sprintf(`select type from bookrzn.Users where Login='%s' AND Password = '%s';`, login, password)) //,

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var access string

	for rows.Next() {
		err := rows.Scan(
			&access)
		if err != nil {
			panic(err)
		}
	}

	return access
}
