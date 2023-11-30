package connector

import (
	"backend/config"
	"backend/models"
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

func (conn *Connector) DataUserFromToken(token string) models.Users {
	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Users WHERE token = '%s' ;`, token))
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var user models.Users

	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Login,
			&user.Password,
			&user.Type,
			&user.Token,
			&user.Name,
			&user.Family,
			&user.Phone,
			&user.Email,
		)

		if err != nil {
			continue
		}

	}

	return user
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

func (conn *Connector) TargetCardsFromListOrders(token string) []models.TargetCard {
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
		temp_target_cards_all = append(temp_target_cards_all, conn.TargetCardFromTargetHash(hash))
	}

	for _, card := range temp_target_cards_all {
		card.Count = mapa_target_hash_count[card.TargetHash]
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
		card.Summa = float64(fc * fp)
		main_target_cards_all = append(main_target_cards_all, card)
	}

	return main_target_cards_all
}

func (conn *Connector) TargetCardsFromListOrdersCMS() []models.TargetCard {
	main_target_cards_all := make([]models.TargetCard, 0)
	for _, token := range conn.ListTokenUsersFromOrders() {
		mapa_target_hash_count := make(map[string]string, 0)
		list_target_hash := make([]string, 0)
		temp_target_cards_all := make([]models.TargetCard, 0)
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
			temp_target_cards_all = append(temp_target_cards_all, conn.TargetCardFromTargetHash(hash))
		}

		for _, card := range temp_target_cards_all {
			user, err := conn.FindUserFromToken(token)
			if err != nil {
				panic(err)
			}

			card.CMSNameOrders = token
			card.CMSPhoneOrders = user.Phone
			card.CMSEmailOrders = user.Email
			// card.CMSPriceAllOrders = user.

			card.Count = mapa_target_hash_count[card.TargetHash]
			card.Price = strings.ReplaceAll(card.Price, ",", ".")
			// card.Title =
			// fc, err := strconv.ParseFloat(card.Count, 64)
			// if err != nil {
			// 	panic(err)
			// }
			// fp, err := strconv.ParseFloat(card.Price, 64)
			// if err != nil {
			// 	panic(err)
			// }
			// card.Price // float64(fc * fp)
			main_target_cards_all = append(main_target_cards_all, card)
		}
	}
	return main_target_cards_all
}

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
			fmt.Println(err)
			continue
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
				fmt.Println(err)
				continue
			}
			fmt.Printf("%+v\n\n", card)
			favcards = append(favcards, card)
		}
	}
	fmt.Println(len(favcards), "favcard")

	return favcards
}

func (conn *Connector) SaveTargetInOrders(token, target_hash, count string) {
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

func (conn *Connector) SaveTargetInFavorites(token, targethash, count string) {
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
			fmt.Println(err)
			continue
		}
	}

	return access
}
