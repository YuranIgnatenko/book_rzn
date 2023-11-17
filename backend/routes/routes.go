package routes

import (
	"backend/auth"
	"backend/connector"
	"backend/datatemp"
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
	// fmt.Println(rout)
	return &rout
}

func (rout *Rout) OpenHtmlAbout(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "about.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlBlog(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "blog.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlCart(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "cart.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlContacts(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "contacts.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlDelivery(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "delivery.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlExchange(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "exchange.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlHome(w http.ResponseWriter, r *http.Request) {
	res := rout.Auth.GetCookieUser(w, r)
	rout.DataTemp.IsLogin = res
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "home.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlLogout(w http.ResponseWriter, r *http.Request) {
	// res := rout.Auth.GetCookieUser(w, r)
	rout.DeleteCookie(w, r)
	rout.DataTemp.IsLogin = false
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "login.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlNew(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "new.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlPayment(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "payment.html")
	tmpl.Execute(w, rout.DataTemp)
}
func (rout *Rout) OpenHtmlCollectschool(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(token)

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
	fmt.Println(token, access)
	switch access {
	case "admin":
		rout.DataTemp.IsLogin = true
		rout.Connector.ReSaveCookieDB(login, password, token)
		rout.Auth.SetCookieAdmin(w, r, token)
		// http.Redirect(w, r, rout.DataTemp.Ip+rout.DataTemp.Split_ip_port+rout.DataTemp.Port+"/cms", http.StatusSeeOther)
		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "cms.html")
		tmpl.Execute(w, rout.DataTemp)
		return
	case "user":
		rout.DataTemp.IsLogin = true
		rout.Connector.ReSaveCookieDB(login, password, token)
		rout.Auth.SetCookieUser(w, r, token)
		// http.Redirect(w, r, rout.DataTemp.Ip+rout.DataTemp.Split_ip_port+rout.DataTemp.Port+"/home", http.StatusSeeOther)
		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "home.html")
		tmpl.Execute(w, rout.DataTemp)
		return
	default:
		// http.Redirect(w, r, rout.DataTemp.Bd_admin_list+"/login", http.StatusSeeOther)
		rout.DataTemp.IsLogin = false
		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "login.html")
		tmpl.Execute(w, rout.DataTemp)
		return
	}
}

func (rout *Rout) OpenHtmlSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	request := r.FormValue("search")

	// token, err := r.Cookie("token")
	// if err != nil {
	// 	return
	// }

	// fmt.Println(r.Cookies())
	// rout.DataTemp.FavoritesCards = rout.GetListFavorites(token.Value)

	rout.DataTemp.SearchTarget = rout.Connector.SearchTargetList(request)
	fmt.Println("len favorites cards::::", rout.DataTemp.SearchTarget)

	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "search.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlFavorites(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	token, err := r.Cookie("token")
	if err != nil {
		return
	}

	// fmt.Println(r.Cookies())
	rout.DataTemp.FavoritesCards = rout.GetListFavorites(token.Value)
	fmt.Println("len favorites cards::::", rout.DataTemp.FavoritesCards)

	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "favorites.html")
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
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "sales.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlProsv(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(w, r)
	// rout.ProsvCards = parsing.ParsingService.ReadFromCsv()
	rout.DataTemp.TargetCards = rout.TargetCards
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "prosv.html")
	tmpl.Execute(w, rout.DataTemp)
}
func (rout *Rout) OpenHtmlAgat(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "agat.html")
	tmpl.Execute(w, rout.DataTemp)
}
func (rout *Rout) OpenHtmlStronikum(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "stronikum.html")
	tmpl.Execute(w, rout.DataTemp)
}
func (rout *Rout) OpenHtmlNaura(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "naura.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtml804(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "804.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlProfile(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL.Path)
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
		// target_hash := strings.Join(url[1:], "//")
		data := r.URL.Path
		data = strings.ReplaceAll(data, ":/", "://")
		fmt.Println("start url add favorites:::", data)

		data = strings.ReplaceAll(data, `"`, "")
		// data = strings.ReplaceAll(data, "'", "")
		target_hash := strings.ReplaceAll(data, "/add_favorites/", "")
		// target_hash = strings.TrimSpace(target_hash)
		// target_hash = strings.ReplaceAll(target_hash, "/", "")
		// target_hash = strings.ReplaceAll(target_hash, "\\", "")

		fmt.Println("add_favorite::::[", target_hash, "]")
		// count := rout.CountRows("bookrzn.Favorites")
		rout.SaveTargetFavorites(token.Value, string(target_hash))
		// fmt.Println(r.Cookies())
		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "prosv.html")
		tmpl.Execute(w, rout.DataTemp)
	case "delete_favorites":
	case "add_orders":
	case "delete_orders":
	default:
	}

	// if string(url[1]) == "add_favorites" {
	// 	order_id := url[2]
	// 	fmt.Println("to be saved ++++++", path, token, order_id)
	// 	count := rout.CountRows("bookrzn.Favorites")
	// 	rout.SaveTargetFavorites(token.Value, string(order_id), string(count))
	// 	fmt.Println(r.Cookies())
	// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "prosv.html")
	// 	tmpl.Execute(w, rout.DataTemp)
	// }

	// if string(url[1]) == "delete_favorites" {
	// 	order_id := url[2]
	// 	fmt.Println("to be deleted -----", path, token, order_id)
	// 	// rout.SaveTarget(token.Value, string(order_id))
	// 	err := rout.DeleteTargetFavorites(token.Value, string(order_id))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	rout.DataTemp.FavoritesCards = rout.GetListFavorites(token.Value)

	// 	fmt.Println(r.Cookies())
	// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "favorites.html")
	// 	tmpl.Execute(w, rout.DataTemp)
	// }

}
