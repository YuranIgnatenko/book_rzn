package connector

import (
	"backend/models"
	"database/sql"
	"errors"
	"fmt"
)

type TableUsers struct {
	DB *sql.DB
}

func NewTableUsers() *TableOrders {
	return &TableOrders{}
}

func (t_users *TableUsers) DataUserFromToken(token string) models.Users {
	rows, err := t_users.DB.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Users WHERE token = '%s' ;`, token))
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
			panic(err)
		}

	}

	return user
}

func (t_users *TableUsers) GetNameLoginFromToken(token string) string {
	rows, err := t_users.DB.Query(fmt.Sprintf(`SELECT login FROM bookrzn.Users WHERE token = '%s' ;`, token)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var login_name string

	for rows.Next() {
		err := rows.Scan(&login_name)

		if err != nil {
			panic(err)
		}

	}

	if login_name == "" {
		return ""
	}
	return login_name
}

func (t_users *TableUsers) AddUser(Login, Password, Type, Token, Name, Family, Phone, Email string) {
	rows, err := t_users.DB.Query(
		fmt.Sprintf(`INSERT INTO bookrzn.Users (login,password,type,token,name,family,phone,email) 
		VALUES ('%s','%s', '%s','%s','%s','%s',' %s', '%s');`,
			Login, Password, Type, Token, Name, Family, Phone, Email))

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
			panic(err)
		}
		users = append(users, u)
	}
}

func (t_users *TableUsers) FindUserFromToken(token string) (models.Users, error) {
	rows, err := t_users.DB.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Users WHERE token = '%s';`, token))
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
			panic(err)
		}
		users = append(users, u)
	}
	if len(users) > 0 {
		return users[0], nil
	} else {
		return models.Users{}, errors.New("error not found user")
	}
}

func (t_users *TableUsers) FindUserFromLoginPassword(Login, Password string) (models.Users, error) {
	rows, err := t_users.DB.Query(fmt.Sprintf(`select * from bookrzn.Users where Login = '%s' AND Password = '%s';`, Login, Password)) //,
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
			panic(err)
		}
	}

	return users, err
}

func (t_users *TableUsers) GetTokenUser(Login, Password string) string {
	rows, err := t_users.DB.Query(fmt.Sprintf(`SELECT token FROM bookrzn.Users WHERE Login='%s' AND Password = '%s';`, Login, Password)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var token string

	for rows.Next() {
		err := rows.Scan(
			&token)
		if err != nil {
			panic(err)
		}

	}

	return token
}

func (t_users *TableUsers) GetAccessUser(Login, Password string) string {
	rows, err := t_users.DB.Query(fmt.Sprintf(`SELECT type FROM bookrzn.Users WHERE login='%s' AND password = '%s';`, Login, Password)) //,
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
