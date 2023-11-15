package connector

import (
	"backend/models"
	"database/sql"
	"fmt"

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

func (conn *Connector) AddUser(login, password, name, phone, email, tp string) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db
	defer db.Close()
	token := "010101"
	rows, err := conn.Db.Query(`INSERT bookrzn.users(Login, Password, Name, Phone, Email, Type, Token) 
								VALUES ('%s', '%s','%s', '%s', '%s', '%s');`,
		login, password, name, phone, email, tp, token)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	users := []models.Users{}

	for rows.Next() {
		u := models.Users{}
		err := rows.Scan(&u.Login, &u.Password, &u.Phone, &u.Email, &u.Type, &u.Token)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	for _, u := range users {
		fmt.Println(u.Login, u.Password, u.Phone, u.Phone)
	}
}

func (conn *Connector) FindUser() {

}
