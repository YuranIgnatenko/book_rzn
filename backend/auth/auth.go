package auth

import (
	"backend/bd"
	"backend/config"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var Config = config.NewConfiguration()

func SetCookieUser(w http.ResponseWriter, r *http.Request, token string) {
	// Initialize a new cookie containing the string "Hello world!" and some
	// non-default attributes.
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	// Use the http.SetCookie() function to send the cookie to the client.
	// Behind the scenes this adds a `Set-Cookie` header to the response
	// containing the necessary cookie data.
	http.SetCookie(w, &cookie)

	// Write a HTTP response as normal.
	// w.Write([]byte("cookie set!"))
}

func GetCookieUser(w http.ResponseWriter, r *http.Request) bool {
	// Retrieve the cookie from the request using its name (which in our case is
	// "exampleCookie"). If no matching cookie is found, this will return a
	// http.ErrNoCookie error. We check for this, and return a 400 Bad Request
	// response to the client.

	// cookie, err := r.Cookie("exampleCookie")
	_, err := r.Cookie("token")

	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			// http.Error(w, "cookie not found", http.StatusBadRequest)
			http.Redirect(w, r, Config.Ip+Config.Split_ip_port+Config.Port+"/404", http.StatusSeeOther)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return false
	} else {
		return true
	}

	// Echo out the cookie value in the response body.
	// w.Write([]byte(cookie.Value))
}

func SetCookieAdmin(w http.ResponseWriter, r *http.Request, token string) {
	// Initialize a new cookie containing the string "Hello world!" and some
	// non-default attributes.
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	// Use the http.SetCookie() function to send the cookie to the client.
	// Behind the scenes this adds a `Set-Cookie` header to the response
	// containing the necessary cookie data.
	http.SetCookie(w, &cookie)

	// Write a HTTP response as normal.
	// w.Write([]byte("cookie set!"))
}

func GetCookieAdmin(w http.ResponseWriter, r *http.Request) bool {
	// Retrieve the cookie from the request using its name (which in our case is
	// "exampleCookie"). If no matching cookie is found, this will return a
	// http.ErrNoCookie error. We check for this, and return a 400 Bad Request
	// response to the client.
	_, err := r.Cookie("token")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			// http.Error(w, "cookie not found", http.StatusBadRequest)
			http.Redirect(w, r, Config.Ip+Config.Split_ip_port+Config.Port+"/404", http.StatusSeeOther)

		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return false
	}

	// Echo out the cookie value in the response body.
	// w.Write([]byte(cookie.Value))
	return true
}

func CheckLoginUser(login, password string) (bool, string) {
	data_rows := bd.ReadUsersData()
	for _, row := range data_rows {
		if login == row[0] {
			fmt.Println(password, row[1])
			if password == row[1] {
				if row[2] == "token" {
					return true, row[3]
				}
			}
		}
	}
	return false, "error"
}

func CheckAdmin(login, password string) (bool, string) {
	data_rows := bd.ReadAdminData()
	// fmt.Println("data row::", data_rows)
	for _, row := range data_rows {
		// fmt.Println("row::", row)
		if login == row[0] {
			if password == row[1] {
				if row[2] == "token" {
					return true, row[3]
				}
			}
		}
	}
	return false, "error"
}
