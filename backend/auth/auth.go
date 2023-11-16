package auth

import (
	"backend/config"
	"backend/connector"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
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

func (a *Auth) DeleteCookie(w http.ResponseWriter, r *http.Request) {

	cookie := http.Cookie{
		Name:  "token",
		Value: "",
		Path:  "/",
		// MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

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
func (a *Auth) CreateUser(login, password, name, family, phone, email string) string {
	// user, err := a.FindUserFromLoginPassword(login, password)
	// if err != nil {
	// 	panic(err)
	// }
	a.AddUser(login, password, "user", a.NewToken(), name, family, phone, email)
	return a.NewToken()
}

func (a *Auth) VerifyLogin(login, password string) (string, string) {
	user, err := a.FindUserFromLoginPassword(login, password)
	if err != nil {
		panic(err)
	}
	access := a.GetAccessUser(login, password)
	fmt.Println(user.Token, access, "verifyyyyy")
	return a.NewToken(), access
}

// todo: add func
func (a *Auth) NewToken() string {
	rand.Seed(time.Now().UnixNano())
	token := rand.Intn(999999999999)
	fmt.Println(token)
	return fmt.Sprintf("%d", token)
}
