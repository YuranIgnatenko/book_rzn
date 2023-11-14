package auth

import (
	"backend/bd"
	"backend/config"
	"errors"
	"log"
	"net/http"
)

type Auth struct {
	Config config.Configuration
	MaxAge int //3600
	Bd     bd.Bd
}

func (a *Auth) SetCookieUser(w http.ResponseWriter, r *http.Request, token string) {

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   a.MaxAge,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

}

func (a *Auth) GetCookieUser(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("token")

	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Redirect(w, r, a.Config.Ip+a.Config.Split_ip_port+a.Config.Port+"/404", http.StatusSeeOther)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return false
	} else {
		return true
	}

}

func (a *Auth) SetCookieAdmin(w http.ResponseWriter, r *http.Request, token string) {

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

func (a *Auth) GetCookieAdmin(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("token")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Redirect(w, r, a.Config.Ip+a.Config.Split_ip_port+a.Config.Port+"/404", http.StatusSeeOther)

		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return false
	}

	return true
}

func (a *Auth) CheckLoginUser(login, password string) (bool, string) {
	data_rows := a.Bd.ReadUsersData()
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

func (a *Auth) CheckAdmin(login, password string) (bool, string) {
	data_rows := a.Bd.ReadAdminData()
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
