package connector

import (
	"backend/config"
	"backend/models"
	"backend/tools"
	"database/sql"
	"fmt"
	"strconv"

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

	tc_all := make([]models.TargetCard, 0)
	for _, hash := range list {
		tc_all = append(tc_all, conn.GetTarget(hash))
	}

	return tc_all
}

func (conn *Connector) GetListOrdersCMS() []models.OrdersCMS {
	var (
		orders_cms             = make([]models.OrdersCMS, 0)
		mapa_token_target_hash = make(map[string][]string, 0)
		mapa_token_user_data   = make(map[string]map[string]string, 0) // {"token1234":{"name":"abc", "phone":"+79008009080"}}
		mapa_token_count_all   = make(map[string]string, 0)
		mapa_token_price_all   = make(map[string]string, 0)
		sum_count              = 0
		sum_price              = 0
	)

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
		tc, err := strconv.Atoi(temp_count)
		if err != nil {
			tc = 0
		}
		sum_count += tc
		mapa_token_target_hash[temp_token] = append(mapa_token_target_hash[temp_token], temp_target_hash)
		mapa_token_count_all[temp_token] = strconv.Itoa(sum_count)

	}

	for hash, _ := range mapa_token_count_all {
		rows, err = conn.Db.Query(fmt.Sprintf(`SELECT price FROM bookrzn.Targets WHERE target_hash='%s';`, hash))
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			temp_price := ""
			temp_token := ""
			err := rows.Scan(&temp_price)
			if err != nil {
				continue
			}
			tp, err := strconv.Atoi(temp_price)
			if err != nil {
				return nil
			}
			sum_price += tp
			mapa_token_price_all[temp_token] = strconv.Itoa(sum_price)

		}
	}

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

	for token_user, data := range mapa_token_user_data {
		temp_order := models.OrdersCMS{}
		temp_order.Name = data["name"]
		temp_order.Phone = data["phone"]
		temp_order.Email = data["email"]
		temp_order.Token = token_user
		orders_cms = append(orders_cms, temp_order)

	}
	fmt.Println(len(orders_cms), "orders_cms")
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
	var value int
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

func (conn *Connector) SaveTargetFavorites(token, targethash string) {
	fmt.Println("save fav cards:", token, targethash)
	// namedb := "bookrzn.Favorites"
	rows, err := conn.Db.Query(
		fmt.Sprintf(`INSERT bookrzn.Favorites (token,target_hash,count) 
		VALUES ( '%s','%s','%s');`,
			token, targethash, "1")) //,
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
