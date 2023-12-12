package auth

import (
	"backend/config"
	"backend/connector"
	"backend/sender"
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

// func GetNameLogin

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

func (a *Auth) GetCookieClient(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("token")
	if err != nil {
		return false
	}
	token := a.GetCookieTokenValue(w, r)
	if token == "" || token == "0" {
		return false
	}

	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return false
	} else {
		return true
	}

}

func (a *Auth) GetCookieTokenValue(w http.ResponseWriter, r *http.Request) string {
	token, err := r.Cookie("token")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			return ""
		default:
			return ""
		}
	} else {
		return token.Value
	}

}

func (a *Auth) SendFormMailValue(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	email := r.FormValue("Email")
	phone := r.FormValue("Phone")
	message := r.FormValue("Message")
	sender.Send_mail("New message from bookrzn", fmt.Sprintf("%v\n%v\n%v\n%v\n", name, email, phone, message))

}

func (a *Auth) DeleteCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    "0",
		Path:     "/",
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
			http.Redirect(w, r, "/404", http.StatusSeeOther)
		default:
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return false
	}

	return true
}
func (a *Auth) CreateUser(login, password, name, family, phone, email string) string {
	a.TableUsers.AddUser(login, password, "user", a.NewToken(), name, family, phone, email)
	return a.NewToken()
}

func (a *Auth) VerifyLogin(login, password string) (string, string) {
	_, err := a.TableUsers.FindUserFromLoginPassword(login, password)
	if err != nil {
		panic(err)
	}
	access := a.TableUsers.GetAccessUser(login, password)
	return a.NewToken(), access
}

// todo: add func
func (a *Auth) NewToken() string {
	rand.Seed(time.Now().UnixNano())
	token := rand.Intn(999999999999)
	return fmt.Sprintf("%d", token)
}
