package connector

var insert_user = `INSERT Users(Id, Login, Password, Name, Phone, Email, Type, Token) 
VALUES ( '%s', '%s', '%s','%s', '%s', '%s', '%s');`

var (
	sql_create_users = `CREATE TABLE Users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    login VARCHAR(30) NOT NULL,
    password VARCHAR(30) NOT NULL,
	type VARCHAR(30) NOT NULL,
	token VARCHAR(30) NOT NULL,
    name VARCHAR(30) ,
	family VARCHAR(30),
	phone VARCHAR(30),
	email VARCHAR(30));`

	sql_create_favorites = `CREATE TABLE Favorites (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
	token VARCHAR(30) NOT NULL,
	target_hash VARCHAR(200) NOT NULL,
	count VARCHAR(30),
	datetime DATETIME);`

	sql_create_orders = `CREATE TABLE Orders (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
	token VARCHAR(30) NOT NULL,
	target_hash VARCHAR(200) NOT NULL,
	count VARCHAR(30),
	datetime DATETIME);`

	sql_create_targets = `CREATE TABLE Targets (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
	target_hash VARCHAR(200) NOT NULL,
	autor VARCHAR(30) NOT NULL,
	title VARCHAR(30) NOT NULL,
	price VARCHAR(30) NOT NULL,
	image VARCHAR(100) NOT NULL,
	comment VARCHAR(500));`
)
