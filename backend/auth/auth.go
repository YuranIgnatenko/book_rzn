package auth

import (
	"backend/bd"
	"backend/config"
	"errors"
	"log"
	"net/http"
)

var Config = config.NewConfiguration()

func SetCookieUser(w http.ResponseWriter, r *http.Request, token string) {

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

}

func GetCookieUser(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("token")

	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Redirect(w, r, Config.Ip+Config.Split_ip_port+Config.Port+"/404", http.StatusSeeOther)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return false
	} else {
		return true
	}

}

func SetCookieAdmin(w http.ResponseWriter, r *http.Request, token string) {

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

}

func GetCookieAdmin(w http.ResponseWriter, r *http.Request) bool {

	_, err := r.Cookie("token")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Redirect(w, r, Config.Ip+Config.Split_ip_port+Config.Port+"/404", http.StatusSeeOther)

		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return false
	}

	return true
}

func CheckLoginUser(login, password string) (bool, string) {
	data_rows := bd.ReadUsersData()
	for _, row := range data_rows {
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

func CheckAdmin(login, password string) (bool, string) {
	data_rows := bd.ReadAdminData()
	for _, row := range data_rows {
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
