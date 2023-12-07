package routes

import (
	"backend/auth"
	"backend/config"
	"backend/connector"
	"backend/datatemp"
	"backend/parsing"
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
	rout.DataTemp.NameLogin = rout.TableUsers.GetNameLoginFromToken(tokenValue)

	path_url := NewPathUrlArgs(r.URL.Path)

	fmt.Println("url:::::::", path_url.ArgRow, path_url.ArgCase)

	switch path_url.ArgCase {
	// добавление в избранное
	case "add_favorites":
		target_hash := path_url.Arg1
		target_count := path_url.Arg2
		fmt.Println("arg1, arg2 ADD:target_hash, target_count:", target_hash, target_count)
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
		target_count := rout.TableOrders.GetCountFromIdOrderHash(target_id_order)
		fmt.Println(target_count)

		rout.DataTemp.TargetCards = rout.TableOrdersHistory.TargetCardsFromListOrdersHistory(tokenValue)
		for _, t_hash := range rout.DataTemp.TargetCards {
			fmt.Println("tokenValue, t_hash.TargetHash, target_count, target_id_order", tokenValue, t_hash.TargetHash, target_count, target_id_order)
			rout.TableOrdersHistory.SaveTargetInOrdersHistory(tokenValue, t_hash.TargetHash, target_count, target_id_order)
		}
		rout.DataTemp.ListOrdersTargetCard = rout.ListOrdersFromTargetCards(rout.DataTemp.TargetCards)
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
	case "confirm_table_orders_history":

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

	// фильтр для товара
	case "book_new":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_new")
		rout.SetHTML(w, "book_prosv.html")

	// фильтр для товара
	case "book_sh_middle":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_sh_middle")
		rout.SetHTML(w, "book_prosv.html")

	// фильтр для товара
	case "book_do_sh":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_do_sh")
		rout.SetHTML(w, "book_prosv.html")

	// фильтр для товара
	case "book_1_4":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_1_4")
		rout.SetHTML(w, "book_prosv.html")

	// фильтр для товара
	case "book_5_9":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_5_9")
		rout.SetHTML(w, "book_prosv.html")

	// фильтр для товара
	case "book_10_11":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_10_11")
		rout.SetHTML(w, "book_prosv.html")

	// фильтр для товара
	case "book_ovz":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_ovz")
		rout.SetHTML(w, "book_prosv.html")

	// фильтр для товара
	case "book_digit_books":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_digit_books")
		rout.SetHTML(w, "book_prosv.html")

	// фильтр для товара
	case "book_artistic":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "prosv_artistic")
		rout.SetHTML(w, "book_prosv.html")

	// фильтр для товара
	case "sh_table":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "sh_table")
		rout.SetHTML(w, "mebel.html")

	// фильтр для товара
	case "sh_chair":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "sh_chair")
		rout.SetHTML(w, "mebel.html")

	// фильтр для товара
	case "office_table":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "office_table")
		rout.SetHTML(w, "mebel.html")

	// фильтр для товара
	case "office_boxing":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "office_boxing")
		rout.SetHTML(w, "mebel.html")

	// фильтр для товара
	case "new_basic":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "new_basic")
		rout.SetHTML(w, "mebel.html")

	// фильтр для товара
	case "new_table":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "new_table")
		rout.SetHTML(w, "mebel.html")

	// фильтр для товара
	case "new_boxing":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "new_boxing")
		rout.SetHTML(w, "mebel.html")

	// фильтр для товара
	case "sh_minitable":
		rout.DataTemp.TargetCards = rout.FilterCards(rout.TargetAll, "sh_minitable")
		rout.SetHTML(w, "mebel.html")

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
		rout.SetHTML(w, "404.html")

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

			http.Redirect(w, r, "/404", http.StatusPermanentRedirect)
			return
		}

	// страница заглушка при ошибке
	case "404":
		rout.SetHTML(w, "404.html")

	// выход с аккаунта и перенаправление на страницу входа
	case "out":
		rout.DataTemp.IsLogin = false
		rout.DataTemp.NameLogin = ""
		rout.DeleteCookie(w, r)
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)

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
