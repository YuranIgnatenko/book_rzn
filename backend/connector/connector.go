package connector

import (
	"backend/config"
	"backend/models"
	"backend/tools"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Connector struct {
	config.Configuration
	Db *sql.DB
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
	//
	return &conn

}

func (conn *Connector) GetNameLoginFromToken(token string) string {
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db

	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT login FROM bookrzn.Users WHERE token = '%s' ;`, token)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var login_name string

	for rows.Next() {
		err := rows.Scan(&login_name)

		if err != nil {
			fmt.Println(err)
			continue
		}

	}

	if login_name == "" {
		return ""
	}
	return login_name
}

func (conn *Connector) SearchTargetList(request string) []models.TargetCard {
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
	}
	conn.Db = db

	var target_search []models.TargetCard

	rows, err := conn.Db.Query(`SELECT * FROM bookrzn.Orders WHERE target_hash LIKE '%` + request + `%' ;`)

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
			fmt.Println(err)
			continue
		}
		target_search = append(target_search, ts)
	}
	return target_search
}

func (conn *Connector) ReSaveCookieDB(login, password, token string) {
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
	}
	conn.Db = db

	old_token := conn.GetTokenUser(login, password)
	new_token := token

	rows, err := conn.Db.Query(fmt.Sprintf(`UPDATE bookrzn.Users SET token = REPLACE(token, '%s', '%s');`, old_token, new_token))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	rows, err = conn.Db.Query(
		fmt.Sprintf(`UPDATE bookrzn.Favorites SET token = REPLACE(token, '%s', '%s');`,
			old_token, token))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

func (conn *Connector) GetListOrders(token string) []models.TargetCard {
	mapa_target_hash_count := make(map[string]string, 0)
	list_target_hash := make([]string, 0)
	temp_target_cards_all := make([]models.TargetCard, 0)
	main_target_cards_all := make([]models.TargetCard, 0)
	temp_token, temp_target_hash, temp_count := "", "", ""

	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT token,target_hash,count FROM bookrzn.Orders WHERE token='%s';`, token)) //,
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
		list_target_hash = append(list_target_hash, temp_target_hash)
		mapa_target_hash_count[temp_target_hash] = temp_count
	}

	for _, hash := range list_target_hash {
		temp_target_cards_all = append(temp_target_cards_all, conn.GetTarget(hash))
	}

	for _, card := range temp_target_cards_all {
		card.Count = mapa_target_hash_count[card.TargetHash]
		main_target_cards_all = append(main_target_cards_all, card)
	}

	return main_target_cards_all
}

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

func (conn *Connector) GetListOrdersCMS() []models.TargetCard {
	var (
		orders_cms             = make([]models.TargetCard, 0)          // card orders list for cms
		mapa_token_target_hash = make(map[string][]string, 0)          // {"token_user":{target_hash1, target_hash2, target_hash3 ....}}
		mapa_token_user_data   = make(map[string]map[string]string, 0) // {"token1234":{"name":"abc", "phone":"+79008009080"}}
		mapa_token_count_all   = make(map[string]string, 0)            // {"token_user":count_all}
		mapa_token_price_all   = make(map[string]string, 0)            // {"token_user":price_all}
	)

	// -------------------------------------------
	// mapa_token_count_all   (COUNT ALL)
	rows, err := conn.Db.Query(`SELECT token,target_hash,count FROM bookrzn.Orders;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		temp_token := ""
		temp_target_hash := ""
		temp_count := ""

		err := rows.Scan(&temp_token, &temp_target_hash, &temp_count)
		if err != nil {
			continue
		}
		tc, err := strconv.Atoi(strings.TrimSpace(temp_count))
		if err != nil {
			tc = 1
		}

		//----------------------------------------------
		//mapa_token_target_hash TOKEN -- > TARGET hash

		mapa_token_target_hash[temp_token] = append(mapa_token_target_hash[temp_token], temp_target_hash)

		var old_count int

		if strings.TrimSpace(mapa_token_count_all[temp_token]) == "" {
			err = nil
			old_count = 0
		} else {
			oc, err := strconv.Atoi(strings.TrimSpace(mapa_token_count_all[temp_token]))
			if err != nil {
				panic(err)
			}
			old_count = oc
		}
		mapa_token_count_all[temp_token] = strconv.Itoa(old_count + tc)

	}

	// ---------------------------------------------------------------------------------------
	//

	rows, err = conn.Db.Query(`SELECT token, name, phone, email  FROM bookrzn.Users;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		temp_token := ""
		temp_name := ""
		temp_phone := ""
		temp_email := ""
		err := rows.Scan(&temp_token, &temp_name, &temp_phone, &temp_email)
		if err != nil {
			continue
		}

		mapa_token_user_data[temp_token] = map[string]string{
			"name":  temp_name,
			"phone": temp_phone,
			"email": temp_email}
	}

	rows, err = conn.Db.Query(`SELECT target_hash,price FROM bookrzn.Targets;`) // target_hash --> price one
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	data_target_prices := make(map[string]float64, 0) //have hash-->price

	for rows.Next() {

		temp_target_hash := ""
		temp_price := ""

		err := rows.Scan(&temp_target_hash, &temp_price)
		if err != nil {
			continue
		}

		temp_price = strings.ReplaceAll(temp_price, "руб", "")
		temp_price = strings.ReplaceAll(temp_price, "₽", "")
		temp_price = strings.ReplaceAll(temp_price, " ", "")
		temp_price = strings.TrimSpace(temp_price)

		t := ""
		for _, sym := range temp_price {
			if string(sym) == "\u00a0" {
				continue
			}
			t += string(sym)
		}

		tp, err := strconv.ParseFloat(t, 64)
		if err != nil {
			tp = float64(0.0)
		}

		data_target_prices[temp_target_hash] = tp
	}

	summa_prices := float64(0.0)

	for token_user, _ := range mapa_token_user_data {
		for _, t_hash := range mapa_token_target_hash[token_user] {
			summa_prices += data_target_prices[t_hash]
		}
		mapa_token_price_all[token_user] = fmt.Sprintf("%v", summa_prices)
		summa_prices = 0.0
	}

	for token_user, data := range mapa_token_user_data {
		// mapa_token_target_hash[token_user]
		temp_order := models.TargetCard{}
		temp_order.CMSName = data["name"]
		temp_order.CMSPhone = data["phone"]
		temp_order.CMSEmail = data["email"]
		temp_order.CMSToken = token_user
		temp_order.CMSTargetsHash = conn.GetTargetHashStringList(token_user)
		temp_order.CMSCountAll = mapa_token_count_all[token_user]
		temp_order.CMSPriceAll = mapa_token_price_all[token_user]

		orders_cms = append(orders_cms, temp_order)
	}
	return orders_cms
}

func (conn *Connector) GetListFavorites(token string) []models.TargetCard {

	list_targets_hash := []string{}

	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT target_hash FROM bookrzn.Favorites WHERE token='%s';`, token)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var favcards []models.TargetCard
	for rows.Next() {
		var target_hash string
		rows.Scan(&target_hash)

		if err != nil {
			fmt.Println(err)
			continue
		}
		list_targets_hash = append(list_targets_hash, target_hash)
	}

	for _, hash := range list_targets_hash {
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

			if err != nil {
				fmt.Println(err)
				continue
			}
			favcards = append(favcards, card)
		}
	}
	fmt.Println(len(favcards), "favcard")

	return favcards
}

func (conn *Connector) GetListOrdersRow(token string) []models.OrdersRows {
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db

	list_targets_hash := []string{}

	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT target_hash FROM bookrzn.Orders WHERE token='%s';`, token))
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var ordcards []models.OrdersRows
	for rows.Next() {

		var target_hash string
		rows.Scan(&target_hash)

		if err != nil {
			fmt.Println(err)
			continue
		}
		list_targets_hash = append(list_targets_hash, target_hash)
	}

	for _, hash := range list_targets_hash {
		rows, err = conn.Db.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Targets WHERE target_hash='%s';`, hash))
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			card := models.OrdersRows{}
			err := rows.Scan(
				&card.Id,

				&card.Price,
				&card.Link,
				&card.Comment,
			)

			if err != nil {
				fmt.Println(err)
				continue
			}
			ordcards = append(ordcards, card)
		}
	}
	return ordcards
}

func (conn *Connector) CountRows(namebd string) int {
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db

	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT COUNT(*) FROM %s;`, namebd)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var count string

	for rows.Next() {
		err := rows.Scan(&count)

		if err != nil {
			fmt.Println(err)
			continue
		}

	}
	c, err := strconv.Atoi(count)
	if err != nil {
		panic(err)

	}
	return c
}

func (conn *Connector) GetNewRandomNumberFastOrder() string {
	return tools.NewRandomTokenFastOrder()
}

func (conn *Connector) GetNewNumberFastOrder() string {
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
	}
	conn.Db = db

	// namedb := "bookrzn.Favorites"
	rows, err := conn.Db.Query(`SELECT MAX(Id) FROM bookrzn.FastOrders;`)
	if err != nil {
		panic(err)
	}
	var value float64
	rows.Scan(&value)

	value += 1
	// обязательно иначе привысит лимит подключений и будет сбой
	defer rows.Close()
	return fmt.Sprintf("%d", value)

}

func (conn *Connector) GetFastOrderList() []models.DataFastOrderOne {
	var rows *sql.Rows
	var err error

	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
	}
	conn.Db = db

	dt := make([]models.DataFastOrderOne, 0)

	rows, err = conn.Db.Query(`SELECT * FROM bookrzn.FastOrders ;`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		dtfast := models.DataFastOrderOne{}
		temp := ""

		err := rows.Scan(
			&temp,
			&dtfast.Name,
			&dtfast.Phone,
			&dtfast.Email,
			&dtfast.Target,
			&dtfast.Count,
			&dtfast.Token,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		dt = append(dt, dtfast)
	}

	// обязательно иначе привысит лимит подключений и будет сбой
	defer rows.Close()
	return dt
}

func (conn *Connector) SaveTargetOrders(token, target_hash, count string) {
	var rows *sql.Rows
	var err error

	rows, err = conn.Db.Query(
		fmt.Sprintf(`INSERT bookrzn.Orders (token, target_hash, count) 
		VALUES ( '%s','%s','%s');`,
			token, target_hash, count))
	if err != nil {
		panic(err)
	}
	// обязательно иначе привысит лимит подключений и будет сбой
	defer rows.Close()
}

func (conn *Connector) SaveTargetFavorites(token, targethash, count string) {
	fmt.Println("save fav cards:", token, targethash, count)
	// namedb := "bookrzn.Favorites"
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

func (conn *Connector) DeleteTargetFavorites(token, targethash string) error {

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
			fmt.Println(err)
			continue
		}
	}

	return access
}
