package routes

import (
	"backend/auth"
	"backend/config"
	"backend/connector"
	"backend/datatemp"
	"backend/models"
	"backend/parsing"
	"backend/sender"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type Rout struct {
	auth.Auth
	connector.Connector
	config.Configuration
	datatemp.DataTemp
	parsing.ParsingService
	CacheMapaTokenSearch map[string]string
	CacheMapaTokenFilter map[string]string
}

func NewRout(a auth.Auth, c config.Configuration, conn connector.Connector, dt datatemp.DataTemp, ps parsing.ParsingService) *Rout {
	rout := Rout{
		Auth:                 a,
		Configuration:        c,
		Connector:            conn,
		DataTemp:             dt,
		ParsingService:       ps,
		CacheMapaTokenSearch: make(map[string]string, 0),
		CacheMapaTokenFilter: make(map[string]string, 0),
	}
	rout.DataTemp.PageTarget.PageSize = 20
	rout.DataTemp.PageTarget.PageNext = 2
	rout.DataTemp.PageTarget.PageNow = 1
	rout.DataTemp.PageTarget.PagePrev = 1
	rout.DataTemp.PageTarget.PageDataAll = rout.ListTargetCardCache
	// rout.DataTemp.PageTarget.PageTotal = len(rout.DataTemp.PageTarget.PageDataAll)

	rout.DataTemp.PageTarget.PageData = rout.DataTemp.PageTarget.GetPage("", 1)
	// rout.DataTemp.OrdersCardsCms = rout.GetListOrdersCMS()

	return &rout
}

func (rout *Rout) DownloadFile(w http.ResponseWriter, r *http.Request) {
	filePath := "static/804.pdf"
	w.Header().Set("Content-Disposition", "attachment; filename=static/804.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, filePath)
	http.Redirect(w, r, "/home", http.StatusPermanentRedirect) // Отправляем файл клиенту
}

func (rout *Rout) ServerSwitchRout(w http.ResponseWriter, r *http.Request) {
	isFindCookie := rout.Auth.GetCookieClient(w, r)
	rout.DataTemp.IsLogin = isFindCookie
	var TokenValue = "host"

	if isFindCookie {
		TokenValue := rout.GetCookieTokenValue(w, r)
		switch rout.TableUsers.GetTypeFromToken(TokenValue) {
		case "admin":
			rout.DataTemp.IsAdmin = true
		default:
			rout.DataTemp.IsAdmin = false
		}
		rout.DataTemp.NameLogin = rout.TableUsers.GetNameLoginFromToken(TokenValue)
	} else {
		rout.DataTemp.IsAdmin = false
		rout.DataTemp.NameLogin = "Гость"
		// http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		// 	// rout.SetHTML(w, "login.html")
	}

	str_url := NewPathUrlArgs(r.URL.Path)

	switch str_url.ArgCase {

	// страница Домашняя
	case "home":
		// path_menu := str_url.Arg1
		// rout.DataTemp.MenuCards
		rout.SetHTML(w, "home.html")
	case
		"new_basic",
		"new_table",
		"new_boxing",
		"sh_table",
		"sh_chair",
		"office_table",
		"office_boxing",
		"sh_minitable":
		var num_page = 1
		if str_url.Arg1 == "" {
			num_page = 1
		} else {
			str_url.Arg1 = strings.ReplaceAll(str_url.Arg1, "?", "")
			num, err := strconv.Atoi(str_url.Arg1)
			if err != nil {
				num_page = 1
			} else {
				num_page = num
			}

		}
		str_url.ArgCase = "nau804"
		rout.DataTemp.PageTarget.PageDataAll = rout.FilterCards(rout.TargetAll, str_url.ArgCase)
		rout.DataTemp.PageTarget.SortedBySwitch(rout.CacheMapaTokenFilter[TokenValue])
		rout.DataTemp.PageTarget.PageData = rout.DataTemp.PageTarget.GetPage(str_url.ArgCase, num_page)

		rout.SetHTML(w, "books.html")

	case
		"str_top",
		"str_psiholog",
		"str_do_sh_3_4",
		"str_do_sh_4_5",
		"str_do_sh_5_6",
		"str_do_sh_6_7",
		"str_sh_started",
		"str_phisic",
		"str_himiya",
		"str_biologiya",
		"str_litra",
		"str_ru_lang",
		"str_other_lang",
		"str_history",
		"str_geograph",
		"str_math",
		"str_info",
		"str_obg",
		"str_eco",
		"str_izo",
		"str_music",
		"str_tehno",
		"str_posters",
		"naura":
		var num_page = 1
		if str_url.Arg1 == "" {
			num_page = 1
		} else {
			str_url.Arg1 = strings.ReplaceAll(str_url.Arg1, "?", "")
			num, err := strconv.Atoi(str_url.Arg1)
			if err != nil {
				num_page = 1
			} else {
				num_page = num
			}

		}
		rout.DataTemp.PageTarget.PageDataAll = rout.FilterCards(rout.TargetAll, str_url.ArgCase)
		rout.DataTemp.PageTarget.PageData = rout.DataTemp.PageTarget.GetPage(str_url.ArgCase, num_page)

		rout.SetHTML(w, "books.html")

	case
		"book_new",
		"book_sh_middle",
		"book_do_sh",
		"book_1_4",
		"book_5_9",
		"book_10_11",
		"book_ovz",
		"book_actistic",
		"book_digit_books":
		var num_page = 1
		if str_url.Arg1 == "" {
			num_page = 1
		} else {
			str_url.Arg1 = strings.ReplaceAll(str_url.Arg1, "?", "")
			num, err := strconv.Atoi(str_url.Arg1)
			if err != nil {
				num_page = 1
			} else {
				num_page = num
			}

		}
		rout.DataTemp.PageTarget.PageDataAll = rout.FilterCards(rout.TargetAll, str_url.ArgCase)
		rout.DataTemp.PageTarget.PageData = rout.DataTemp.PageTarget.GetPage(str_url.ArgCase, num_page)

		rout.SetHTML(w, "books.html")

	case "search":
		res := r.FormValue("search")

		if res == "" {
			res = fmt.Sprint(rout.CacheMapaTokenSearch[TokenValue])
		} else {
			rout.CacheMapaTokenSearch[TokenValue] = res
		}

		rout.DataTemp.LastValueSearch = res

		var num_page = 1
		if str_url.Arg1 != "" {

			str_url.Arg1 = strings.ReplaceAll(str_url.Arg1, "?", "")
			num, err := strconv.Atoi(str_url.Arg1)
			if err != nil {
				num_page = 1
			} else {
				num_page = num
			}
		}
		// rout.PageTarget.LastSearch = res
		rout.DataTemp.PageTarget.PageDataAll = rout.FilterSearch(rout.TargetAll, res)
		rout.PageTarget.PageTotal = len(rout.DataTemp.PageTarget.PageDataAll)
		rout.DataTemp.PageTarget.PageData = rout.DataTemp.PageTarget.GetPage(str_url.ArgCase, num_page)

		rout.SetHTML(w, "search.html")

	case "set_filter_book":
		fmt.Println(">>>>>>>>>>>>>")
		fmt.Println(r.URL.Path)
		// pa := GetFilterUrlArgs(r.URL.Path)
		rout.CacheMapaTokenFilter[TokenValue] = str_url.Arg1
		// rout.DataTemp.PageTarget.PageData = rout.TableTargets.GetListTargetsFromTokenHistoryStatusOFF(TokenValue)
		// fmt.Println(pa.Arg1)
		http.Redirect(w, r, "/book_1_4", http.StatusPermanentRedirect)

	case "home_news":
		http.Redirect(w, r, "/home", http.StatusPermanentRedirect)

	case "home_contacts_address":
		rout.SetHTML(w, "home_contacts_address.html")

	case "home_docs_info":
		rout.SetHTML(w, "home_docs_info.html")

	case "orders_history_cancel_orders":
		// rout.DataTemp.TargetCards = rout.TableTargets.GetListTargetsFromTokenHistoryStatusOFF(TokenValue)
		rout.DataTemp.PageTarget.PageData = rout.TableTargets.GetListTargetsFromTokenHistoryStatusOFF(TokenValue)
		rout.DataTemp.ListOrdersTargetCard = rout.ListOrdersFromTargetCards(rout.DataTemp.PageTarget.PageData)
		rout.SetHTML(w, "orders_history_cancel_orders.html")

	case "cancel_orders_history":
		target_id_order := str_url.Arg1
		rout.TableOrdersHistory.CancelTableOrdersHistory(TokenValue, target_id_order)
		sender.Send_mail("ЗАКАЗ", "Отменён заказ на сайте")
		fmt.Println("Email sended CANCELED ORDER !")
		http.Redirect(w, r, "/orders_history", http.StatusPermanentRedirect)
	// для копирования из истории заказов в избранное
	// case "move_fav_table_orders":

	case "home_804":
		rout.DownloadFile(w, r)
		rout.SetHTML(w, "home_804.html")

	case "home_vk":
		http.Redirect(w, r, "https://vk.com/magazin_rzn", http.StatusPermanentRedirect)

	case "load_logo":
		r.ParseMultipartForm(10 << 20) // Парсим форму с максимальным размером 10MB

		file, _, err := r.FormFile("file")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		f, err := os.OpenFile("static/logo_new.jpg", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		io.Copy(f, file)

	case "load_banner":
		r.ParseMultipartForm(10 << 20) // Парсим форму с максимальным размером 10MB

		file, _, err := r.FormFile("file")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		f, err := os.OpenFile("static/banner.jpg", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		io.Copy(f, file)

	// добавление в избранное
	case "card_view":
		target_hash := str_url.Arg1
		// rout.DataTemp.TargetCards = []models.TargetCard{rout.TableTargets.GetTargetsCardsFromHash(target_hash)}
		rout.DataTemp.PageTarget.PageData = []models.TargetCard{rout.TableTargets.GetTargetsCardsFromHash(target_hash)}
		rout.SetHTML(w, "card_view.html")

	// добавление в избранное
	case "add_favorites":
		if !rout.DataTemp.IsLogin {
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}
		target_hash := str_url.Arg1
		target_count := str_url.Arg2
		rout.SaveTargetInFavorites(TokenValue, string(target_hash), target_count)
		http.Redirect(w, r, "/favorites", http.StatusPermanentRedirect)

	// удаление из избранного
	case "delete_favorites":
		target_hash := str_url.Arg1
		rout.DeleteTargetFromFavorites(TokenValue, string(target_hash))
		http.Redirect(w, r, "/favorites", http.StatusPermanentRedirect)

	// добавление в черновые таблицы заказов
	case "add_orders":
		target_hash := str_url.Arg1
		target_count := str_url.Arg2
		target_id_order := str_url.Arg3
		rout.SaveTargetInOrders(TokenValue, string(target_hash), target_count, target_id_order)
		http.Redirect(w, r, "/favorites", http.StatusPermanentRedirect)

	// удаление таблицы из черновых заказов
	case "delete_table_orders":
		target_id_order := str_url.Arg1
		rout.TableOrders.DeleteTableOrders(rout.GetCookieTokenValue(w, r), target_id_order)
		http.Redirect(w, r, "/orders", http.StatusPermanentRedirect)

	// сахранить\перенос в историю заказов таблицы
	case "confirm_table_orders":
		target_id_order := str_url.Arg1

		rout.TableOrdersHistory.SaveTargetInOrdersHistory(TokenValue, target_id_order)

		sender.Send_mail("ЗАКАЗ", "Выполнен заказ на сайте")
		http.Redirect(w, r, "/orders_history", http.StatusPermanentRedirect)

	//страница ИСТОРИЯ активных и прошлых заказов
	case "orders_history":
		// rout.DataTemp.TargetCards = rout.TableTargets.GetListTargetsFromTokenHistoryStatusON(TokenValue)
		rout.DataTemp.PageTarget.PageData = rout.TableTargets.GetListTargetsFromTokenHistoryStatusON(TokenValue)
		rout.DataTemp.ListOrdersTargetCard = rout.ListOrdersFromTargetCards(rout.DataTemp.PageTarget.PageData)
		rout.SetHTML(w, "orders_history.html")

	case "delete_table_orders_history":
		target_id_order := str_url.Arg1
		rout.DeleteTableOrdersHistory(rout.GetCookieTokenValue(w, r), target_id_order)
		// rout.DataTemp.TargetCards = rout.TableTargets.GetListTargetsFromTokenHistoryStatusOFF(TokenValue)
		rout.DataTemp.PageTarget.PageData = rout.TableTargets.GetListTargetsFromTokenHistoryStatusOFF(TokenValue)
		rout.DataTemp.ListOrdersTargetCard = rout.ListOrdersFromTargetCards(rout.DataTemp.PageTarget.PageData)
		http.Redirect(w, r, "/orders_history_cancel_orders", http.StatusPermanentRedirect)

	case "edit_table_orders":
		rout.DataTemp.PageTarget.PageData = rout.TableTargets.GetListTargetsFromToken(TokenValue)
		rout.DataTemp.ListOrdersTargetCard = rout.ListOrdersFromTargetCards(rout.DataTemp.PageTarget.PageData)
		rout.SetHTML(w, "orders_editor.html")

	case "delete_record_orders":
		target_hash := str_url.Arg1
		rout.TableOrders.DeleteRecordOrders(rout.GetCookieTokenValue(w, r), target_hash)
		rout.DataTemp.PageTarget.PageData = rout.TableTargets.GetListTargetsFromToken(TokenValue)
		rout.DataTemp.ListOrdersTargetCard = rout.ListOrdersFromTargetCards(rout.DataTemp.PageTarget.PageData)
		http.Redirect(w, r, "/orders", http.StatusPermanentRedirect)
		// rout.SetHTML(w, "orders.html")

	// страницы черновых заказов- составление
	case "orders":
		rout.DataTemp.PageTarget.PageData = rout.TableTargets.GetListTargetsFromToken(TokenValue)
		rout.DataTemp.ListOrdersTargetCard = rout.ListOrdersFromTargetCards(rout.DataTemp.PageTarget.PageData)
		rout.SetHTML(w, "orders.html")

	// страница с избранными карточками товаров
	case "favorites":
		rout.DataTemp.FavoritesCards = rout.TargetCardsFromListFavorites(TokenValue)
		rout.SetHTML(w, "favorites.html")

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
		if rout.DataTemp.IsLogin {
			if rout.GetCookieAdmin(w, r) {
				// rout.DataTemp.TargetCards = rout.TargetCardsFromListOrdersCMS()
				rout.SetHTML(w, "cms.html")
				return
			}
		}
		rout.SetHTML(w, "404.html")

	// case "cms_view_order":

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

			http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
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
		rout.DataTemp.NameLogin = "Гость"
		rout.DeleteCookie(w, r)
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)

	default:
		fmt.Printf("\n[ ERR PATH ] -- [ ERR PATH ] -- [ %v ]", r.URL.Path)
		// http.Redirect(w, r, "/404", http.StatusPermanentRedirect)

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
