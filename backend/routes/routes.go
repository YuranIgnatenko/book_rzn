package routes

import (
	"backend/auth"
	"backend/connector"
	"backend/datatemp"
	"backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type Rout struct {
	auth.Auth
	connector.Connector
	datatemp.DataTemp
}

func NewRout(a auth.Auth, conn connector.Connector, dt datatemp.DataTemp) *Rout {
	rout := Rout{
		Auth:      a,
		Connector: conn,
		DataTemp:  dt,
	}

	return &rout
}

func (rout *Rout) OpenHtmlFastOrder(w http.ResponseWriter, r *http.Request) {
	rout.NumberFastOrder = rout.GetNewRandomNumberFastOrder()
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "fast_order.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlFastOrderSave(w http.ResponseWriter, r *http.Request) {
	var data models.DataFastOrder

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	rout.Connector.SaveTargetFastOrders(data)

	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "home.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlAbout(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "about.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlBlog(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "blog.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlCart(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "cart.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlContacts(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "contacts.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlDelivery(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "delivery.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlExchange(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "exchange.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlHome(w http.ResponseWriter, r *http.Request) {

	token := rout.GetCookieToken(w, r)
	_, err := rout.Connector.FindUserFromToken(token)
	if err != nil {
		rout.DataTemp.IsLogin = false
	} else {
		rout.DataTemp.IsLogin = true
	}

	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "home.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlLogout(w http.ResponseWriter, r *http.Request) {

	rout.DeleteCookie(w, r)
	// rout.SetCookieUser(w,r,)
	rout.DataTemp.IsLogin = false
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "login.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlNew(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "new.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlPayment(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "payment.html")
	tmpl.Execute(w, rout.DataTemp)
}
func (rout *Rout) OpenHtmlCollectschool(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "collect-school.html")
	tmpl.Execute(w, rout.DataTemp)
}
func (rout *Rout) OpenHtmlLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "login.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlRegistry(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	name := r.FormValue("name")
	family := r.FormValue("family")
	phone := r.FormValue("phone")
	email := r.FormValue("email")

	token := rout.CreateUser(login, password, name, family, phone, email)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)

	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "registration.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlCms(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieAdmin(w, r)
	if res {
		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "cms.html")
		tmpl.Execute(w, rout.DataTemp)
	} else {
		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "404.html")
		tmpl.Execute(w, rout.DataTemp)
	}

}

func (rout *Rout) OpenHtmlLoginCheck(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	token, access := rout.VerifyLogin(login, password)
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)

	switch access {
	case "admin":
		fmt.Println("admin -- ok")
		rout.DataTemp.IsLogin = true
		rout.Connector.ReSaveCookieDB(login, password, token)
		rout.Auth.SetCookieAdmin(w, r, token)

		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "cms.html")
		tmpl.Execute(w, rout.DataTemp)
		return
	case "user":
		rout.DataTemp.IsLogin = true
		rout.Connector.ReSaveCookieDB(login, password, token)
		rout.Auth.SetCookieUser(w, r, token)

		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "home.html")
		tmpl.Execute(w, rout.DataTemp)
		return
	default:

		rout.DataTemp.IsLogin = false
		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "login.html")
		tmpl.Execute(w, rout.DataTemp)
		return
	}
}

func (rout *Rout) OpenHtmlSearch(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)

	request := r.FormValue("search")

	rout.DataTemp.SearchTarget = rout.Connector.SearchTargetList(request)

	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "search.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlFavorites(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		return
	}

	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token.Value)
	rout.DataTemp.FavoritesCards = rout.GetListFavorites(token.Value)

	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "favorites.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlOrders(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		return
	}

	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token.Value)
	rout.DataTemp.OrdersRows = rout.GetListOrdersRow(token.Value)

	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "orders.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtml404(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "404.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlCreateUser(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "login.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlSales(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "sales.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlProsv(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)

	rout.DataTemp.TargetCards = rout.TargetCards
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "prosv.html")
	tmpl.Execute(w, rout.DataTemp)
}
func (rout *Rout) OpenHtmlAgat(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "agat.html")
	tmpl.Execute(w, rout.DataTemp)
}
func (rout *Rout) OpenHtmlStronikum(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "stronikum.html")
	tmpl.Execute(w, rout.DataTemp)
}
func (rout *Rout) OpenHtmlNaura(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "naura.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtml804(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	token := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "804.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlProfile(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	t := rout.GetCookieToken(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(t)

	url := strings.Split(r.URL.Path, "/")
	if url[0] == "" {
		url = url[1:]
	}

	token, err := r.Cookie("token")
	if err != nil {
		return
	}

	switch url[0] {
	case "add_favorites":
		data := r.URL.Path
		data = strings.ReplaceAll(data, ":/", "://")
		data = strings.ReplaceAll(data, `"`, "")
		target_hash := strings.ReplaceAll(data, "/add_favorites/", "")
		rout.SaveTargetFavorites(token.Value, string(target_hash))

		http.Redirect(w, r, "/prosv", http.StatusPermanentRedirect)

	case "delete_favorites":
		data := r.URL.Path
		data = strings.ReplaceAll(data, ":/", "://")
		data = strings.ReplaceAll(data, `"`, "")
		target_hash := strings.ReplaceAll(data, "/delete_favorites/", "")
		rout.DeleteTargetFavorites(token.Value, string(target_hash))

		http.Redirect(w, r, "/favorites", http.StatusPermanentRedirect)

	case "add_orders":
	case "delete_orders":
	default:
	}
}
