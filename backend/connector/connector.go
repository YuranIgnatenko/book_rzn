package connector

import (
	"backend/models"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "bookrzn"
	password = "book1995"
	hostname = "127.0.0.1:3306"
	dbname   = "bookrzn"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func Connect() {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	res, err := db.Exec(`SHOW DATABASES;`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

var sql_create_users = `CREATE TABLE Users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    login VARCHAR(30) NOT NULL,
    password VARCHAR(30) NOT NULL,
	type VARCHAR(30) NOT NULL,
	token VARCHAR(30) NOT NULL,
    name VARCHAR(30) ,
	family VARCHAR(30),
	phone VARCHAR(30),
	email VARCHAR(30));`

var sql_create_favorites = `CREATE TABLE Favorites (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
	token VARCHAR(30) NOT NULL,
	target_hash VARCHAR(200) NOT NULL,
	count VARCHAR(30),
	datetime DATETIME
);`

var sql_create_orders = `CREATE TABLE Orders (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
	token VARCHAR(30) NOT NULL,
	target_hash VARCHAR(200) NOT NULL,
	count VARCHAR(30),
	datetime DATETIME);`

var sql_create_targets = `CREATE TABLE Targets (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
	target_hash VARCHAR(200) NOT NULL,
	autor VARCHAR(30) NOT NULL,
	title VARCHAR(30) NOT NULL,
	price VARCHAR(30) NOT NULL,
	image VARCHAR(100) NOT NULL,
	comment VARCHAR(500));`

// type (
// 	TableUsers     []models.Users
// 	TableFavorites []models.Favorites
// 	TableOrders    []models.Orders
// 	TableTargets   []models.Targets
// )

type Connector struct {
	username    string
	password    string
	hostname    string
	dbname      string
	BdUsers     string
	BdFavorites string
	BdOrders    string
	BdTargets   string
	Db          *sql.DB
}

func NewConnector(user, pswd, host, dbname string) *Connector {
	conn := Connector{
		username:    user,
		password:    pswd,
		hostname:    host,
		dbname:      dbname,
		BdUsers:     "Users",
		BdFavorites: "Favorites",
		BdOrders:    "Orders",
		BdTargets:   "Targets",
	}

	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db
	defer db.Close()
	return &conn

}

func (conn *Connector) AddUser(Login, Password, Type, Token, Name, Family, Phone, Email string) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db
	defer db.Close()
	// token := "010101"
	namedb := "bookrzn.Users"
	rows, err := conn.Db.Query(
		fmt.Sprintf(`INSERT %s (id,login,password,type,token,name,family,phone,email) 
		VALUES ('%v','%s', '%s','%s','%s','%s',' %s', '%s', '%s','%s');`,
			namedb, conn.CountRows(namedb)+1, Login, Password, Type, Token, Name, Family, Phone, Email)) //,
	//login, password, name, phone, email, tp, token)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	users := []models.Users{}

	for rows.Next() {
		u := models.Users{}
		err := rows.Scan(
			&u.Id,
			&u.Login,
			&u.Password,
			&u.Type,
			&u.Token,
			&u.Name,
			&u.Family,
			&u.Phone,
			&u.Email)

		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	fmt.Println(len(users))
	for _, u := range users {
		fmt.Println(u.Login, u.Password, u.Phone, u.Phone)
	}
}

func (conn *Connector) FindUserFromToken(token string) []models.Users {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db
	defer db.Close()
	rows, err := conn.Db.Query(`select * from bookrzn.Users where Token = '` + token + `';`) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	users := []models.Users{}

	for rows.Next() {
		u := models.Users{}
		err := rows.Scan(
			&u.Id,
			&u.Login,
			&u.Password,
			&u.Type,
			&u.Token,
			&u.Name,
			&u.Family,
			&u.Phone,
			&u.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}

	return users
}

func (conn *Connector) FindUserFromLoginPassword(Login, Password string) (models.Users, error) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
		return models.Users{}, err

	}
	conn.Db = db
	defer db.Close()
	rows, err := conn.Db.Query(fmt.Sprintf(`select * from bookrzn.Users where Login = '%s' AND Password = '%s';`, Login, Password)) //,
	if err != nil {
		return models.Users{}, err
	}

	defer rows.Close()

	users := models.Users{}

	for rows.Next() {
		u := models.Users{}
		err := rows.Scan(
			&u.Id,
			&u.Login,
			&u.Password,
			&u.Type,
			&u.Token,
			&u.Name,
			&u.Family,
			&u.Phone,
			&u.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return users, err
}

func (conn *Connector) GetTokenUser(Login, Password string) string {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db
	defer db.Close()
	rows, err := conn.Db.Query(fmt.Sprintf(`select token from bookrzn.Users where Login='%s' AND Password = '%s';`, Login, Password)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var token string

	for rows.Next() {
		err := rows.Scan(
			&token)
		if err != nil {
			fmt.Println(err)
			continue
		}

	}

	return token
}

func (conn *Connector) GetAccessUser(Login, Password string) string {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db
	defer db.Close()
	rows, err := conn.Db.Query(fmt.Sprintf(`select type from bookrzn.Users where Login='%s' AND Password = '%s';`, Login, Password)) //,
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

func (conn *Connector) GetListFavorites(token string) []models.FavoritesCards {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db
	defer db.Close()
	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT TargetHash FROM bookrzn.Favorites;`)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var favcards []models.FavoritesCards
	for rows.Next() {
		card := models.FavoritesCards{}
		err := rows.Scan(
			&card.Autor,
			&card.Title,
			&card.Price,
			&card.Link,
			&card.Id)

		if err != nil {
			fmt.Println(err)
			continue
		}
		favcards = append(favcards, card)
	}

	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT TargetHash FROM bookrzn.Favorites;`)) //,
	if err != nil {
		panic(err)
	}


	var list_targethash []string
	var th string
	for rows.Next() {
		
		err := rows.Scan(&th)

		if err != nil {
			fmt.Println(err)
			continue
		}
		list_targethash = append(list_targethash, th)
	}


	return c
}

func (conn *Connector) CountRows(namebd string) int {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db
	defer db.Close()
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
		fmt.Println(count)
	}
	c, err := strconv.Atoi(count)
	if err != nil {
		panic(err)

	}
	return c
}

func (conn *Connector) SaveTargetFavorites(token, targethash, count string) {
	dt := time.Now()
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
	}
	conn.Db = db
	defer db.Close()
	namedb := "bookrzn.Favorites"
	rows, err := conn.Db.Query(
		fmt.Sprintf(`INSERT bookrzn.Favorites (id,token,target_hash,count, datetime) 
		VALUES ('%s', '%s','%s','%s','%s',' %s');`,
			conn.CountRows(namedb)+1, token, targethash, count, dt)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

}

func (conn *Connector) DeleteTargetFavorites(token, targethash string) error {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
	}
	conn.Db = db
	defer db.Close()
	rows, err := conn.Db.Query(
		fmt.Sprintf(`DELETE FROM bookrzn.Favorites
		WHERE Token='%s' AND TargetHash='%s';`, token, targethash))
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	return err

}

func (conn *Connector) AccessLogin(login, password string) string {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db
	defer db.Close()
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
