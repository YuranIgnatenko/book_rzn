package connector

import (
	"backend/models"
	"database/sql"
	"fmt"
)

func (conn *Connector) AddUser(Login, Password, Type, Token, Name, Family, Phone, Email string) {
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db
	defer db.Close()
	namedb := "bookrzn.Users"
	rows, err := conn.Db.Query(
		fmt.Sprintf(`INSERT INTO %s (login,password,type,token,name,family,phone,email) 
		VALUES ('%s','%s', '%s','%s','%s','%s',' %s', '%s');`,
			namedb, Login, Password, Type, Token, Name, Family, Phone, Email)) //,
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
	// fmt.Println(len(users))
	for _, u := range users {
		fmt.Println(u.Login, u.Password, u.Phone, u.Phone)
	}
}

func (conn *Connector) FindUserFromToken(token string) []models.Users {
	db, err := sql.Open("mysql", conn.dsn())
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
	db, err := sql.Open("mysql", conn.dsn())
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
	db, err := sql.Open("mysql", conn.dsn())
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
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db
	defer db.Close()
	rows, err := conn.Db.Query(fmt.Sprintf(`select type from bookrzn.Users where login='%s' AND password = '%s';`, Login, Password)) //,
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
