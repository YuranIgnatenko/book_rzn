package routes

import (
	"backend/auth"
	"backend/config"
	"net/http"
	"text/template"
)

type Router struct{
	Auth auth.Auth
	Config config.Configuration
}


func (router *Router) OpenHtmlAbout(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "about.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtmlBlog(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "blog.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtmlCart(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "cart.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtmlContacts(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "contacts.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtmlDelivery(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "delivery.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtmlExchange(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "exchange.html")
	tmpl.Execute(w, router.Config)
}
func (router *Router) OpenHtmlFavorites(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "cart.html")
	tmpl.Execute(w, router.Config)
}
func (router *Router) OpenHtmlHome(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "home.html")
	tmpl.Execute(w, router.Config)
}
func (router *Router) OpenHtmlNew(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "new.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtmlPayment(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "payment.html")
	tmpl.Execute(w, router.Config)
}
func (router *Router) OpenHtmlCollectschool(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "collect-school.html")
	tmpl.Execute(w, router.Config)
}
func (router *Router) OpenHtmlLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "login.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtmlRegistry(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "registration.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtmlCms(w http.ResponseWriter, r *http.Request) {
	res := router.Auth.GetCookieAdmin(w, r)
	if res {
		tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "cms.html")
		tmpl.Execute(w, router.Config)
	} else {
		tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "404.html")
		tmpl.Execute(w, router.Config)
	}

}

func (router *Router) OpenHtmlProfile(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "home.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtmlLoginCheck(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	is_admin, cookie_admin := router.Auth.CheckAdmin(login, password)
	is_user, cookie_user := router.Auth.CheckLoginUser(login, password)

	if is_admin {
		router.Auth.SetCookieAdmin(w, r, cookie_admin)
		http.Redirect(w, r, router.Config.Ip+router.Config.Split_ip_port+router.Config.Port+"/cms", http.StatusSeeOther)
		return
	}
	if is_user {
		router.Auth.SetCookieUser(w, r, cookie_user)
		http.Redirect(w, r, router.Config.Ip+router.Config.Split_ip_port+router.Config.Port+"/home", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, router.Config.Bd_admin_list+"/login", http.StatusSeeOther)
	return

}

func (router *Router) OpenHtmlBuyOrder(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "buy_order.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtml404(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "404.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtmlCreateUser(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "login.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtmlSales(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "sales.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtmlProsv(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "prosv.html")
	tmpl.Execute(w, router.Config)
}
func (router *Router) OpenHtmlAgat(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "agat.html")
	tmpl.Execute(w, router.Config)
}
func (router *Router) OpenHtmlStronikum(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "stronikum.html")
	tmpl.Execute(w, router.Config)
}
func (router *Router) OpenHtmlNaura(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "naura.html")
	tmpl.Execute(w, router.Config)
}

func (router *Router) OpenHtml804(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(router.Config.Path_prefix + router.Config.Path_frontend + "804.html")
	tmpl.Execute(w, router.Config)
}
