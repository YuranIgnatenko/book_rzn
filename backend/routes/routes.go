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
	fmt.Println(r.Cookies())
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "home.html")
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

func (rout *Rout) OpenHtmlProfile(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "home.html")
	tmpl.Execute(w, rout.DataTemp)
}

func (rout *Rout) OpenHtmlLoginCheck(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	token, access := rout.VerifyLogin(login, password)

	switch access {
	case "admin":
		rout.Auth.SetCookieAdmin(w, r, token)
		http.Redirect(w, r, rout.DataTemp.Ip+rout.DataTemp.Split_ip_port+rout.DataTemp.Port+"/cms", http.StatusSeeOther)
		return
	case "user":
		rout.Auth.SetCookieUser(w, r, token)
		http.Redirect(w, r, rout.DataTemp.Ip+rout.DataTemp.Split_ip_port+rout.DataTemp.Port+"/home", http.StatusSeeOther)
		return
	default:
		http.Redirect(w, r, rout.DataTemp.Bd_admin_list+"/login", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, rout.DataTemp.Bd_admin_list+"/login", http.StatusSeeOther)
	return

}

func (rout *Rout) OpenHtmlAddFavorites(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)
	url := strings.Split(r.URL.Path, "/")
	// if url[0] != "add_favorites" {
	// 	http.Redirect(w, r, rout.DataTemp.Ip+rout.DataTemp.Split_ip_port+rout.DataTemp.Port+"/home", http.StatusSeeOther)
	// }
	token, err := r.Cookie("token")
	if err != nil {
		return
	}
	path := r.URL.Path

	fmt.Println("url:", url)

	if string(url[1]) == "add_favorites" {
		order_id := url[2]
		fmt.Println("to be saved ++++++", path, token, order_id)
		rout.SaveTargetFavorites(token.Value, string(order_id))
		fmt.Println(r.Cookies())
		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "prosv.html")
		tmpl.Execute(w, rout.DataTemp)
	}

	if string(url[1]) == "delete_favorites" {
		order_id := url[2]
		fmt.Println("to be deleted -----", path, token, order_id)
		// rout.SaveTarget(token.Value, string(order_id))
		err := rout.DeleteTargetFavorites(token.Value, string(order_id))
		if err != nil{
			panic(err)
		}
		rout.DataTemp.FavoritesCards = rout.GetListFavorites(token.Value)

		fmt.Println(r.Cookies())
		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "favorites.html")
		tmpl.Execute(w, rout.DataTemp)
	}

}

func (rout *Rout) OpenHtmlFavorites(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	token, err := r.Cookie("token")
	if err != nil {
		return
	}

	fmt.Println(r.Cookies())
	rout.DataTemp.FavoritesCards = rout.FindTarget(token.Value)

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
	fmt.Println(w, r)
	// rout.ProsvCards = parsing.ParsingService.ReadFromCsv()
	rout.DataTemp.ProsvCards = rout.ProsvCards
	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "prosv.html")
	fmt.Println("----------", len(rout.DataTemp.ProsvCards))
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
