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
	"text/template"
)

type Rout struct {
	auth.Auth
	connector.Connector
	config.Configuration
	datatemp.DataTemp
	parsing.ParsingService
	AccertPath map[string]int
}

func NewRout(a auth.Auth, c config.Configuration, conn connector.Connector, dt datatemp.DataTemp, ps parsing.ParsingService) *Rout {
	rout := Rout{
		Auth:           a,
		Configuration:  c,
		Connector:      conn,
		DataTemp:       dt,
		ParsingService: ps,
		AccertPath: map[string]int{
			//книги
			"new_basic":     1,
			"new_table":     1,
			"new_boxing":    1,
			"sh_table":      1,
			"sh_chair":      1,
			"office_table":  1,
			"office_boxing": 1,
			"sh_minitable":  1,

			// 	оборудование
			"str_top":        1,
			"str_psiholog":   1,
			"str_do_sh_3_4":  1,
			"str_do_sh_4_5":  1,
			"str_do_sh_5_6":  1,
			"str_do_sh_6_7":  1,
			"str_sh_started": 1,
			"str_phisic":     1,
			"str_himiya":     1,
			"str_biologiya":  1,
			"str_litra":      1,
			"str_ru_lang":    1,
			"str_other_lang": 1,
			"str_history":    1,
			"str_geograph":   1,
			"str_math":       1,
			"str_info":       1,
			"str_obg":        1,
			"str_eco":        1,
			"str_izo":        1,
			"str_music":      1,
			"str_tehno":      1,
			"str_posters":    1,

			// мебель
			"book_new":         1,
			"book_sh_middle":   1,
			"book_do_sh":       1,
			"book_1_4":         1,
			"book_5_9":         1,
			"book_10_11":       1,
			"book_ovz":         1,
			"book_actistic":    1,
			"book_digit_books": 1,
		},
	}
	rout.DataTemp.TargetCards = rout.ListTargetCardCache
	// rout.DataTemp.OrdersCardsCms = rout.GetListOrdersCMS()

	return &rout
}

// прверяем наличия пути в карте разрешенных в качестве товаров
func (rout *Rout) RangePathsTargetPage(w http.ResponseWriter, path string) bool {
	return rout.AccertPath[path] == 1
}

func (rout *Rout) ServerRoutHtml(w http.ResponseWriter, r *http.Request) {
	isFindCookie := rout.Auth.GetCookieClient(w, r)
	rout.DataTemp.IsLogin = isFindCookie

	tokenValue := rout.GetCookieTokenValue(w, r)
	rout.DataTemp.NameLogin = rout.TableUsers.GetNameLoginFromToken(tokenValue)

	path_url := NewPathUrlArgs(r.URL.Path)

	if rout.RangePathsTargetPage(w, path_url.ArgCase) {

		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, path_url.ArgCase)

		rout.SetHTML(w, "targets.html")
		return
	}

	switch path_url.ArgCase {
	// добавление в избранное
	case "add_favorites":
		target_hash := path_url.Arg1
		target_count := path_url.Arg2
		rout.SaveTargetInFavorites(tokenValue, string(target_hash), target_count)
		http.Redirect(w, r, "/favorites", http.StatusPermanentRedirect)

	// удаление из избранного
	case "delete_favorites":
		target_hash := path_url.Arg1
		rout.DeleteTargetFromFavorites(tokenValue, string(target_hash))
		http.Redirect(w, r, "/favorites", http.StatusPermanentRedirect)

	// добавление в черновые таблицы заказов
	case "add_orders":
		target_hash := path_url.Arg1
		target_count := path_url.Arg2
		target_id_order := path_url.Arg3
		rout.SaveTargetInOrders(tokenValue, string(target_hash), target_count, target_id_order)
		http.Redirect(w, r, "/orders", http.StatusPermanentRedirect)

	// удаление таблицы из черновых заказов
	case "delete_table_orders":
		target_id_order := path_url.Arg1
		rout.TableOrders.DeleteTableOrders(rout.GetCookieTokenValue(w, r), target_id_order)
		http.Redirect(w, r, "/orders", http.StatusPermanentRedirect)

	case "move_fav_table_orders":

	// сахранить\перенос в историю заказов таблицы
	case "confirm_table_orders":
		target_id_order := path_url.Arg1
		rout.TableOrdersHistory.SaveTargetInOrdersHistory(tokenValue, target_id_order)
		sender.Send_mail("ЗАКАЗ", "Выполнен заказ на сайте")
		fmt.Println("Email sended !")
		rout.SetHTML(w, "orders_history.html")

	//страница ИСТОРИЯ активных и прошлых заказов
	case "orders_history":
		rout.DataTemp.TargetCards = rout.TableTargets.GetListTargetsFromTokenHistory(tokenValue)
		rout.DataTemp.ListOrdersTargetCard = rout.ListOrdersFromTargetCards(rout.DataTemp.TargetCards)
		rout.SetHTML(w, "orders_history.html")

	// удаление таблицы из страницы ИСТОРИЯ
	// case "delete_table_orders_history":
	// 	target_id_order := path_url.Arg1
	// 	rout.DeleteTableOrdersHistory(rout.GetCookieTokenValue(w, r), target_id_order)
	// 	rout.DataTemp.TargetCards = rout.TargetCardsFromListOrdersHistory(tokenValue)
	// 	rout.DataTemp.ListOrdersTargetCard = rout.ListOrdersFromTargetCards(rout.DataTemp.TargetCards)
	// 	http.Redirect(w, r, "/orders_history", http.StatusPermanentRedirect)

	case "edit_table_orders_history":
	case "move_fav_table_orders_history":

	// case "edit_orders":
	// 	rout.DataTemp.TargetCards = rout.TargetCardsFromListOrders(tokenValue)
	// 	rout.SetHTML(w, "edit_orders.html")

	// страницы черновых заказов- составление
	case "orders":
		rout.DataTemp.TargetCards = rout.TableTargets.GetListTargetsFromToken(tokenValue)
		rout.DataTemp.ListOrdersTargetCard = rout.ListOrdersFromTargetCards(rout.DataTemp.TargetCards)
		rout.SetHTML(w, "orders.html")

	// страница с избранными карточками товаров
	case "favorites":
		rout.DataTemp.TargetCards = rout.TargetCardsFromListFavorites(tokenValue)
		rout.SetHTML(w, "favorites.html")

	// страница Домашняя
	case "home":
		rout.SetHTML(w, "home.html")

	// страница с поиском товара
	case "search":
		sub := r.FormValue("search")
		rout.DataTemp.TargetCards = rout.FilterSearch(rout.DataTemp.TargetAll, sub)
		rout.SetHTML(w, "search.html")

	// страница быстрых писем\заказов
	case "fast_mail":
		rout.SetHTML(w, "fast_mail.html")

	// приём формы быстрого письма\заказа
	case "fast_mail_receive":
		rout.SendFormMailValue(w, r)
		fmt.Println("Email sended !")
		rout.SetHTML(w, "home.html")

	// панель администратора
	case "cms":
		if isFindCookie {
			if rout.GetCookieAdmin(w, r) {
				// rout.DataTemp.TargetCards = rout.TargetCardsFromListOrdersCMS()
				rout.SetHTML(w, "cms.html")
				return
			}
		}
		// rout.SetHTML(w, "404.html")

	// case "cms_view_order":
	// 	token_user := strings.Split(r.URL.Path, "/")[1]
	// 	fmt.Println(token_user)
	// write html table orders list

	// страница входа с полями логин и пороль
	case "login":
		rout.DataTemp.IsLogin = false
		rout.DataTemp.NameLogin = ""
		rout.DeleteCookie(w, r)
		rout.SetHTML(w, "login.html")

	// страница для регистрации
	case "registration":
		rout.SetHTML(w, "registration.html")

	// приём формы регистрации
	case "create_user":
		login := r.FormValue("login")
		password := r.FormValue("password")
		name := r.FormValue("name")
		family := r.FormValue("family")
		phone := r.FormValue("phone")
		email := r.FormValue("email")

		token := rout.CreateUser(login, password, name, family, phone, email)
		rout.DataTemp.NameLogin = rout.TableUsers.GetNameLoginFromToken(token)

		rout.SetHTML(w, "login.html")

	// прием формы входа и валидация
	case "validation_login":
		login := r.FormValue("login")
		password := r.FormValue("password")

		token, access := rout.VerifyLogin(login, password)

		// определение типа токена админ\пользователь
		switch access {
		case "admin":
			rout.DataTemp.NameLogin = rout.TableUsers.GetNameLoginFromToken(token)
			rout.DataTemp.IsLogin = true
			old_token := rout.TableUsers.GetTokenUser(login, password)
			// rout.Connector.ReSaveCookieDB(login, password, token)
			rout.Auth.SetCookieAdmin(w, r, old_token)

			http.Redirect(w, r, "/cms", http.StatusPermanentRedirect)
			return
		case "user":
			rout.DataTemp.NameLogin = rout.TableUsers.GetNameLoginFromToken(token)
			rout.DataTemp.IsLogin = true
			// rout.Connector.ReSaveCookieDB(login, password, token)
			old_token := rout.TableUsers.GetTokenUser(login, password)
			rout.Auth.SetCookieUser(w, r, old_token)

			http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
			return
		default:
			rout.DataTemp.IsLogin = false
			// http.Redirect(w, r, "/404", http.StatusPermanentRedirect)
			return
		}

	// страница заглушка при ошибке
	// case "404":
	// 	rout.SetHTML(w, "404.html")

	// выход с аккаунта и перенаправление на страницу входа
	case "out":
		rout.DataTemp.IsLogin = false
		rout.DataTemp.NameLogin = ""
		rout.DeleteCookie(w, r)
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)

	default:
		// fmt.Printf("\n\n[%v]\n\n", path_url.ArgCase)
		// fmt.Println("open RANGE TARGET DEFAULT-CASE")

		// } else {
		// 	http.Redirect(w, r, "/404", http.StatusPermanentRedirect)
		// }
	}

}

// установка html страницы по названию файла
func (rout *Rout) SetHTML(w http.ResponseWriter, filename string) {
	tmpl, err := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + filename)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, rout.DataTemp)
}
