package routes

import (
	"backend/auth"
	"backend/config"
	"fmt"
	"net/http"
	"text/template"
)

var Config = config.NewConfiguration()

func OpenHtmlAbout(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "about.html")
	tmpl.Execute(w, Config)
}

func OpenHtmlBlog(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "blog.html")
	tmpl.Execute(w, Config)
}

func OpenHtmlCart(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "cart.html")
	tmpl.Execute(w, Config)
}

func OpenHtmlContacts(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "contacts.html")
	tmpl.Execute(w, Config)
}

func OpenHtmlDelivery(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "delivery.html")
	tmpl.Execute(w, Config)
}

func OpenHtmlExchange(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "exchange.html")
	tmpl.Execute(w, Config)
}
func OpenHtmlFavorites(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "cart.html")
	tmpl.Execute(w, Config)
}
func OpenHtmlHome(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("open home", Config.Path_prefix + Config.Path_frontend)
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "home.html")
	tmpl.Execute(w, Config)
}
func OpenHtmlNew(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "new.html")
	tmpl.Execute(w, Config)
}

func OpenHtmlPayment(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "payment.html")
	tmpl.Execute(w, Config)
}
func OpenHtmlCollectschool(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "collect-school.html")
	tmpl.Execute(w, Config)
}
func OpenHtmlLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "login.html")
	tmpl.Execute(w, Config)
}

func OpenHtmlRegistry(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "registration.html")
	tmpl.Execute(w, Config)
}

func OpenHtmlCms(w http.ResponseWriter, r *http.Request) {
	res := auth.GetCookieAdmin(w, r)
	fmt.Println(res, "res cookie admin -- 404 ??")
	if res {
		tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "cms.html")
		tmpl.Execute(w, Config)
	} else {
		tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "404.html")
		tmpl.Execute(w, Config)
	}

}

func OpenHtmlProfile(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "home.html")
	tmpl.Execute(w, Config)
}

func OpenHtmlLoginCheck(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	fmt.Println(login, password)

	is_admin, cookie_admin := auth.CheckAdmin(login, password)
	is_user, cookie_user := auth.CheckLoginUser(login, password)

	fmt.Println(login, password)
	fmt.Println("user mode:", is_user, "\n", "admin mode:", is_admin)

	if is_admin {
		auth.SetCookieAdmin(w, r, cookie_admin)
		http.Redirect(w, r, Config.Full_url_addr+"/cms", http.StatusSeeOther)
		// OpenHtmlCms(w, r)
		return
	}
	if is_user {
		auth.SetCookieUser(w, r, cookie_user)
		// OpenHtmlProfile(w, r)
		http.Redirect(w, r, Config.Full_url_addr+"/home", http.StatusSeeOther)
		return
	}

	// OpenHtmlLogin(w, r)
	http.Redirect(w, r, Config.Bd_admin_list+"/login", http.StatusSeeOther)
	return

}

func OpenHtmlBuyOrder(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "buy_order.html")
	tmpl.Execute(w, Config)
}

func OpenHtml404(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "404.html")
	tmpl.Execute(w, Config)
}

func OpenHtmlCreateUser(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "login.html")
	tmpl.Execute(w, Config)
}

func OpenHtmlSales(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "sales.html")
	tmpl.Execute(w, Config)
}

func OpenHtmlProsv(w http.ResponseWriter, r *http.Request) {
	
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "prosv.html")
	tmpl.Execute(w, Config)
}
func OpenHtmlAgat(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "agat.html")
	tmpl.Execute(w, Config)
}
func OpenHtmlStronikum(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "stronikum.html")
	tmpl.Execute(w, Config)
}
func OpenHtmlNaura(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(Config.Path_prefix + Config.Path_frontend + "naura.html")
	tmpl.Execute(w, Config)
}
