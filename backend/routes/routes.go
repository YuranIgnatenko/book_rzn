package routes

import (
	"backend/auth"
	"backend/config"
	"backend/connector"
	"backend/datatemp"
	"backend/parsing"
	"backend/sender"
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
	// rout.DataTemp.OrdersCardsCms = rout.GetListOrdersCMS()

	return &rout
}

func (rout *Rout) ServerRoutHtml(w http.ResponseWriter, r *http.Request) {
	isFindCookie := rout.Auth.GetCookieClient(w, r)
	rout.DataTemp.IsLogin = isFindCookie

	tokenValue := rout.GetCookieTokenValue(w, r)
	rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(tokenValue)

	url := strings.Split(r.URL.Path, "/")
	if url[0] == "" {
		url = url[1:]
	}

	switch url[0] {

	case "add_favorites":
		data := r.URL.Path
		data = strings.ReplaceAll(data, ":/", "://")
		data = strings.ReplaceAll(data, `"`, "")
		temp := strings.ReplaceAll(data, "/add_favorites/", "") //res:  1701168638010694862/13
		target_hash := strings.Split(temp, "/")[0]
		target_count := strings.Split(temp, "/")[1]
		if strings.TrimSpace(target_count) == "" || target_count == "0" {
			target_count = "1"
		}
		rout.SaveTargetInFavorites(tokenValue, string(target_hash), target_count)
		http.Redirect(w, r, "/home", http.StatusPermanentRedirect)

	case "delete_favorites":
		data := r.URL.Path
		data = strings.ReplaceAll(data, ":/", "://")
		data = strings.ReplaceAll(data, `"`, "")
		target_hash := strings.ReplaceAll(data, "/delete_favorites/", "")
		rout.DeleteTargetFromFavorites(tokenValue, string(target_hash))

		http.Redirect(w, r, "/favorites", http.StatusPermanentRedirect)

	case "add_orders":
		fmt.Println("add orders")
		data := r.URL.Path
		data = strings.ReplaceAll(data, ":/", "://")
		data = strings.ReplaceAll(data, `"`, "")
		temp := strings.ReplaceAll(data, "/add_orders/", "")
		target_hash := strings.Split(temp, "/")[0]
		target_count := strings.Split(temp, "/")[1]
		if strings.TrimSpace(target_count) == "" || target_count == "0" {
			target_count = "1"
		}
		rout.SaveTargetInOrders(tokenValue, string(target_hash), target_count)

		fmt.Println("Loaded sender")

		user := rout.DataUserFromToken(tokenValue)
		target := rout.TargetCardFromTargetHash(target_hash)

		data_msg := fmt.Sprintf(
			"Создан заказ! \nКонтактное лицо :[( %v ) %v %v]\n\n Связь:[ %v %v ]\n\n Позиция:[ %v %v ] \nКоличество:[ %v ] \nЦена(за 1 экз.):[ %v ].",
			user.Login, user.Name, user.Family,
			user.Email, user.Phone, target.Autor, target.Title, target_count, target.Price)

		sender.Send_mail("Уведомление о заказе",
			fmt.Sprintf("%v\n", data_msg))

		fmt.Println("Email sended !")

		http.Redirect(w, r, "/home", http.StatusPermanentRedirect)

	case "orders":
		rout.DataTemp.TargetCards = rout.TargetCardsFromListOrders(tokenValue)
		rout.SetHTML(w, "orders.html")

	case "favorites":
		rout.DataTemp.TargetCards = rout.TargetCardsFromListFavorites(tokenValue)
		rout.SetHTML(w, "favorites.html")

	case "home":
		rout.SetHTML(w, "home.html")

	case "book_new":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_new")
		rout.SetHTML(w, "book_prosv.html")

	case "book_sh_middle":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_sh_middle")
		rout.SetHTML(w, "book_prosv.html")

	case "book_do_sh":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_do_sh")
		rout.SetHTML(w, "book_prosv.html")

	case "book_1_4":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_1_4")
		rout.SetHTML(w, "book_prosv.html")

	case "book_5_9":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_5_9")
		rout.SetHTML(w, "book_prosv.html")

	case "book_10_11":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_10_11")
		rout.SetHTML(w, "book_prosv.html")

	case "book_ovz":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_ovz")
		rout.SetHTML(w, "book_prosv.html")

	case "book_digit_books":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_digit_books")
		rout.SetHTML(w, "book_prosv.html")

	case "book_artistic":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_artistic")
		rout.SetHTML(w, "book_prosv.html")

	case "sh_table":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "sh_table")
		rout.SetHTML(w, "mebel.html")

	case "sh_chair":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "sh_chair")
		rout.SetHTML(w, "mebel.html")

	case "office_table":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "office_table")
		rout.SetHTML(w, "mebel.html")

	case "office_boxing":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "office_boxing")
		rout.SetHTML(w, "mebel.html")

	case "new_basic":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "new_basic")
		rout.SetHTML(w, "mebel.html")

	case "new_table":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "new_table")
		rout.SetHTML(w, "new_table.html")

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
		rout.SendFormMailValue(w, r)
		fmt.Println("Email sended !")
		rout.SetHTML(w, "home.html")

	case "cms":
		if isFindCookie {
			if rout.GetCookieAdmin(w, r) {
				rout.DataTemp.TargetCards = rout.TargetCardsFromListOrdersCMS()
				rout.SetHTML(w, "cms.html")
				return
			}
		}
		rout.SetHTML(w, "404.html")

	// case "cms_view_order":
	// 	token_user := strings.Split(r.URL.Path, "/")[1]
	// 	fmt.Println(token_user)
	// write html table orders list

	case "login":
		rout.DataTemp.IsLogin = false
		rout.DataTemp.NameLogin = ""
		rout.DeleteCookie(w, r)
		rout.SetHTML(w, "login.html")

	case "registration":
		rout.SetHTML(w, "registration.html")

	case "create_user":
		login := r.FormValue("login")
		password := r.FormValue("password")
		name := r.FormValue("name")
		family := r.FormValue("family")
		phone := r.FormValue("phone")
		email := r.FormValue("email")

		token := rout.CreateUser(login, password, name, family, phone, email)
		rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)

		rout.SetHTML(w, "login.html")

	case "validation_login":
		login := r.FormValue("login")
		password := r.FormValue("password")

		token, access := rout.VerifyLogin(login, password)

		switch access {
		case "admin":
			rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
			rout.DataTemp.IsLogin = true
			old_token := rout.GetTokenUser(login, password)
			// rout.Connector.ReSaveCookieDB(login, password, token)
			rout.Auth.SetCookieAdmin(w, r, old_token)

			http.Redirect(w, r, "/cms", http.StatusPermanentRedirect)
			return
		case "user":
			rout.DataTemp.NameLogin = rout.Connector.GetNameLoginFromToken(token)
			rout.DataTemp.IsLogin = true
			// rout.Connector.ReSaveCookieDB(login, password, token)
			old_token := rout.GetTokenUser(login, password)
			rout.Auth.SetCookieUser(w, r, old_token)

			http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
			return
		default:
			rout.DataTemp.IsLogin = false

			http.Redirect(w, r, "/404", http.StatusPermanentRedirect)
			return
		}

	case "404":
		rout.SetHTML(w, "404.html")

	case "out":
		rout.DataTemp.IsLogin = false
		rout.DataTemp.NameLogin = ""
		rout.DeleteCookie(w, r)
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)

	default:
		fmt.Println("error : DEFAULT CASE::[", url, "]")
	}
}

func (rout *Rout) SetHTML(w http.ResponseWriter, filename string) {
	tmpl, err := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + filename)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, rout.DataTemp)
}
