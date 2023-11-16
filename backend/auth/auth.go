package auth

import (
	"backend/config"
	"backend/connector"
	"errors"
	"log"
	"net/http"
)

type Auth struct {
	config.Configuration
	connector.Connector

	MaxAge int //3600

}

func NewAuth(c config.Configuration, conn connector.Connector) *Auth {
	return &Auth{
		Configuration: c,
		Connector:     conn,
	}
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
			http.Redirect(w, r, a.Ip+a.Split_ip_port+a.Port+"/404", http.StatusSeeOther)
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
		Name:  "token",
		Value: token,
		Path:  "/",
		// MaxAge:   3600,
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
			http.Redirect(w, r, a.Ip+a.Split_ip_port+a.Port+"/404", http.StatusSeeOther)

		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return false
	}

	return true
}

func (a *Auth) VerifyLogin(login, password string) (string, string) {
	user, err := a.FindUserFromLoginPassword(login, password)
	if err != nil {
		panic(err)
	}
	access := a.GetAccessUser(login, password)

	return user.Token, access
}
