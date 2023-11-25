package routes

import (
	"backend/auth"
	"backend/config"
	"backend/connector"
	"backend/datatemp"
	"backend/parsing"
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type Rout struct {
	auth.Auth
	connector.Connector
	config.Configuration
	datatemp.DataTemp
	parsing.ParsingService
}

func NewRout(a auth.Auth, c config.Configuration, conn connector.Connector, dt datatemp.DataTemp) *Rout {
	rout := Rout{
		Auth:           a,
		Configuration:  c,
		Connector:      conn,
		DataTemp:       dt,
		ParsingService: *parsing.NewParsingService(c, conn),
	}
	rout.DataTemp.TargetCards = rout.ListTargetCardCache

	return &rout
}

// func (rout *Rout) OpenHtmlFastOrder(w http.ResponseWriter, r *http.Request) {
// 	rout.NumberFastOrder = rout.GetNewRandomNumberFastOrder()
// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "fast_order.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlFastOrderSave(w http.ResponseWriter, r *http.Request) {
// 	var data models.DataFastOrder

// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&data)
// 	if err != nil {
// 		panic(err)
// 	}

// 	rout.Connector.SaveTargetFastOrders(data)

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "home.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlAbout(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "about.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlBlog(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "blog.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlCart(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "cart.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlContacts(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "contacts.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlDelivery(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "delivery.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlExchange(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "exchange.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlHome(w http.ResponseWriter, r *http.Request) {
// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "home.html")
// 	tmpl.Execute(w, rout.DataTemp)

// }

// func (rout *Rout) OpenHtmlLogout(w http.ResponseWriter, r *http.Request) {
// 	rout.DataTemp.IsLogin = false
// 	rout.DataTemp.NameLogin = "Гость"
// 	rout.DeleteCookie(w, r)
// 	http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
// }

// func (rout *Rout) OpenHtmlNew(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "new.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlPayment(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "payment.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }
// func (rout *Rout) OpenHtmlCollectschool(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "collect-school.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }
// func (rout *Rout) OpenHtmlLogin(w http.ResponseWriter, r *http.Request) {
// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "login.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlRegistry(w http.ResponseWriter, r *http.Request) {
// 	login := r.FormValue("login")
// 	password := r.FormValue("password")
// 	name := r.FormValue("name")
// 	family := r.FormValue("family")
// 	phone := r.FormValue("phone")
// 	email := r.FormValue("email")

// 	token := rout.CreateUser(login, password, name, family, phone, email)
// 	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "registration.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlCms(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieAdmin(w, r)
// 	rout.FastOrdersList = rout.GetFastOrderList()
// 	if res {
// 		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "cms.html")
// 		tmpl.Execute(w, rout.DataTemp)
// 	} else {
// 		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "404.html")
// 		tmpl.Execute(w, rout.DataTemp)
// 	}

// }

// func (rout *Rout) OpenHtmlLoginCheck(w http.ResponseWriter, r *http.Request) {
// 	login := r.FormValue("login")
// 	password := r.FormValue("password")

// 	token, access := rout.VerifyLogin(login, password)
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res
// 	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)

// 	switch access {
// 	case "admin":
// 		rout.DataTemp.IsLogin = true
// 		rout.Connector.ReSaveCookieDB(login, password, token)
// 		rout.Auth.SetCookieAdmin(w, r, token)

// 		http.Redirect(w, r, "/cms", http.StatusPermanentRedirect)
// 		return
// 	case "user":
// 		rout.DataTemp.IsLogin = true
// 		rout.Connector.ReSaveCookieDB(login, password, token)
// 		rout.Auth.SetCookieUser(w, r, token)

// 		http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
// 		return
// 	default:

// 		rout.DataTemp.IsLogin = false
// 		http.Redirect(w, r, "/404", http.StatusPermanentRedirect)
// 		return
// 	}
// }

// func (rout *Rout) OpenHtmlSearch(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	request := r.FormValue("search")

// 	rout.DataTemp.SearchTarget = rout.Connector.SearchTargetList(request)

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "search.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlFavorites(w http.ResponseWriter, r *http.Request) {
// 	token, err := r.Cookie("token")
// 	if err != nil {
// 		return
// 	}

// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res
// 	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token.Value)
// 	rout.DataTemp.FavoritesCards = rout.GetListFavorites(token.Value)

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "favorites.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlOrders(w http.ResponseWriter, r *http.Request) {
// 	token, err := r.Cookie("token")
// 	if err != nil {
// 		return
// 	}

// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res
// 	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token.Value)
// 	rout.DataTemp.OrdersRows = rout.GetListOrdersRow(token.Value)

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "orders.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtml404(w http.ResponseWriter, r *http.Request) {
// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "404.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlCreateUser(w http.ResponseWriter, r *http.Request) {
// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "login.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlSales(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "sales.html")
// 	tmpl.Execute(w, rout.DataTemp)
// // }

// func (rout *Rout) OpenHtmlProsv(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	if res {
// 		token := rout.GetCookieTokenValue(w, r)
// 		rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
// 	}
// 	rout.DataTemp.TargetCards = rout.TargetCards

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "prosv.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlAgat(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "agat.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }
// func (rout *Rout) OpenHtmlStronikum(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "stronikum.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }
// func (rout *Rout) OpenHtmlNaura(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "naura.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtml804(w http.ResponseWriter, r *http.Request) {
// 	res := rout.Auth.GetCookieClient(w, r)
// 	rout.DataTemp.IsLogin = res

// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "804.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

func (rout *Rout) ServerRoutHtml(w http.ResponseWriter, r *http.Request) {
	isFindCookie := rout.Auth.GetCookieClient(w, r)
	rout.DataTemp.IsLogin = isFindCookie

	tokenValue := rout.GetCookieTokenValue(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(tokenValue)

	url := strings.Split(r.URL.Path, "/")
	if url[0] == "" {
		url = url[1:]
	}

	fmt.Println(url, len(url))

	switch url[0] {

	// case "for_school":

	case "add_favorites":
		data := r.URL.Path
		data = strings.ReplaceAll(data, ":/", "://")
		data = strings.ReplaceAll(data, `"`, "")
		target_hash := strings.ReplaceAll(data, "/add_favorites/", "")
		rout.SaveTargetFavorites(tokenValue, string(target_hash))
		http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
		fmt.Println("add favorites")

	case "delete_favorites":
		data := r.URL.Path
		data = strings.ReplaceAll(data, ":/", "://")
		data = strings.ReplaceAll(data, `"`, "")
		target_hash := strings.ReplaceAll(data, "/delete_favorites/", "")
		rout.DeleteTargetFavorites(tokenValue, string(target_hash))

		http.Redirect(w, r, "/favorites", http.StatusPermanentRedirect)

	case "add_orders":
		data := r.URL.Path
		data = strings.ReplaceAll(data, ":/", "://")
		data = strings.ReplaceAll(data, `"`, "")
		target_hash := strings.ReplaceAll(data, "/add_orders/", "")

		count := "1"
		rout.SaveTargetOrders(tokenValue, string(target_hash), count)
		http.Redirect(w, r, "/home", http.StatusPermanentRedirect)

		// case "delete_orders":

	case "orders":
		rout.DataTemp.TargetCards = rout.GetListOrders(tokenValue)
		rout.SetHTML(w, "orders.html")

	case "favorites":
		fmt.Println(tokenValue)
		rout.DataTemp.TargetCards = rout.GetListFavorites(tokenValue)
		fmt.Println("len data /favorites", len(rout.DataTemp.TargetCards))
		rout.SetHTML(w, "favorites.html")

	case "home":
		rout.SetHTML(w, "home.html")

	case "book_prosv":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "book_prosv")
		rout.SetHTML(w, "book_prosv.html")

	case "sh_table":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "sh_table")
		rout.SetHTML(w, "sh_table.html")

	case "sh_chair":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "sh_chair")
		rout.SetHTML(w, "sh_chair.html")

	case "office_table":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "office_table")
		rout.SetHTML(w, "office_table.html")

	case "office_boxing":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "office_boxing")
		rout.SetHTML(w, "office_boxing.html")

	case "new_basic":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "new_basic")
		rout.SetHTML(w, "new_basic.html")

	// case "new_table":
	// 	rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "new_table")
	// 	rout.SetHTML(w, "new_table.html")

	case "new_boxing":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "new_boxing")
		rout.SetHTML(w, "new_boxing.html")

	case "sh_minitable":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "sh_minitable")
		rout.SetHTML(w, "sh_minitable.html")

	case "search":
		sub := r.FormValue("search")
		rout.DataTemp.TargetCards = rout.FilterSearch(rout.DataTemp.TargetAll, sub)
		rout.SetHTML(w, "search.html")

	case "fast_mail":
		rout.SetHTML(w, "fast_mail.html")

	case "fast_mail_receive":
		rout.SetHTML(w, "home.html")

	case "cms":
		rout.SetHTML(w, "cms.html")

	case "login":
		rout.SetHTML(w, "login.html")

	case "registration":
		rout.SetHTML(w, "registration.html")

	case "validation_login":
		login := r.FormValue("login")
		password := r.FormValue("password")

		token, access := rout.VerifyLogin(login, password)
		res := rout.Auth.GetCookieClient(w, r)
		rout.DataTemp.IsLogin = res
		rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)

		switch access {
		case "admin":
			rout.DataTemp.IsLogin = true
			rout.Connector.ReSaveCookieDB(login, password, token)
			rout.Auth.SetCookieAdmin(w, r, token)

			http.Redirect(w, r, "/cms", http.StatusPermanentRedirect)
			return
		case "user":
			rout.DataTemp.IsLogin = true
			rout.Connector.ReSaveCookieDB(login, password, token)
			rout.Auth.SetCookieUser(w, r, token)

			http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
			return
		default:

			rout.DataTemp.IsLogin = false
			http.Redirect(w, r, "/404", http.StatusPermanentRedirect)
			return
		}

	case "404":
		rout.SetHTML(w, "404.html")

	case "logout":
		rout.SetHTML(w, "logout.html")

	default:
		fmt.Println("error : DEFAULT CASE::[", url, "]")
		// http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
	}
}

func (rout *Rout) SetHTML(w http.ResponseWriter, filename string) {
	tmpl, err := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + filename)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, rout.DataTemp)
}

// func (rout *Rout) OpenHtmlTargetCards(w http.ResponseWriter, r *http.Request) {
// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "targetcards.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }

// func (rout *Rout) OpenHtmlForSchool(w http.ResponseWriter, r *http.Request) {
// 	rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetCards, "prosv-book")
// 	fmt.Println(len(rout.DataTemp.TargetCards))
// 	tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "for_school.html")
// 	tmpl.Execute(w, rout.DataTemp)
// }
