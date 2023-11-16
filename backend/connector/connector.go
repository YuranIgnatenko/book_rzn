package connector

import (
	"backend/config"
	"backend/models"
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

func (conn *Connector) ReSaveCookieDB(login, password, token string) {
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
	}
	conn.Db = db

	type User struct {
		login, password, tp, token, name, family, phone, email string
	}
	var user User

	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Users WHERE login='%s' and password='%s';`, login, password)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			&user.login,
			&user.password,
			&user.tp,
			&user.token,
			&user.name,
			&user.family,
			&user.phone,
			&user.email,
		)

		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	rows, err = conn.Db.Query(
		// fmt.Sprintf(`INSERT bookrzn.Users (login,password,type, token,name,family, phone,email)
		fmt.Sprintf(`UPDATE bookrzn.Users SET token='%s' WHERE login='%s' AND password='%s';`,
		token, login, password)) //,
	if err != nil {
		panic(err)
	}
	//
	// обязательно иначе привысит лимит подключений и будет сбой
	defer rows.Close()
}

func (conn *Connector) GetListFavorites(token string) []models.FavoritesCards {
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db

	list_targets_hash := []string{}

	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT target_hash FROM bookrzn.Favorites WHERE token='%s';`, token)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var favcards []models.FavoritesCards
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
			card := models.FavoritesCards{}
			err := rows.Scan(
				&card.Id,
				&card.TargetHash,
				&card.Autor,
				&card.Title,
				&card.Price,
				&card.Link,
				&card.Comment,
			)

			if err != nil {
				fmt.Println(err)
				continue
			}
			favcards = append(favcards, card)
		}
	}
	return favcards
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
		// fmt.Println(count)
	}
	c, err := strconv.Atoi(count)
	if err != nil {
		panic(err)

	}
	return c
}

func (conn *Connector) SaveTargetFavorites(token, targethash string) {
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
	}
	conn.Db = db

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
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
	}
	conn.Db = db

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
