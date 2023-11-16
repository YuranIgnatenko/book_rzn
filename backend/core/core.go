package core

import (
	"backend/auth"
	"backend/config"
	"backend/connector"
	"backend/datatemp"
	"backend/middleware"
	"backend/parsing"
	"backend/routes"
	"fmt"
	"log"
	"net/http"
)

type Core struct {
	*config.Configuration
	*connector.Connector
	*auth.Auth
	*middleware.Middleware
	*parsing.ParsingService
	*routes.Rout
}

func (c *Core) SetHandlers() {
	http.HandleFunc("/about", c.OpenHtmlAbout)
	http.HandleFunc("/blog", c.OpenHtmlBlog)
	http.HandleFunc("/cart", c.CookieUser(http.HandlerFunc(c.OpenHtmlCart)))
	http.HandleFunc("/contacts", c.OpenHtmlContacts)
	http.HandleFunc("/delivery", c.CookieUser(http.HandlerFunc(c.OpenHtmlDelivery)))
	http.HandleFunc("/exchange", c.CookieUser(http.HandlerFunc(c.OpenHtmlExchange)))
	http.HandleFunc("/favorites", c.CookieUser(http.HandlerFunc(c.OpenHtmlFavorites)))
	http.HandleFunc("/home", c.OpenHtmlHome)
	http.HandleFunc("/new", c.CookieUser(http.HandlerFunc(c.OpenHtmlNew)))
	http.HandleFunc("/payment", c.CookieUser(http.HandlerFunc(c.OpenHtmlPayment)))
	http.HandleFunc("/collectschool", c.CookieUser(http.HandlerFunc(c.OpenHtmlCollectschool)))
	http.HandleFunc("/login", c.OpenHtmlLogin)
	http.HandleFunc("/registration", c.OpenHtmlRegistry)
	http.HandleFunc("/cms", c.CookieAdmin(http.HandlerFunc(c.OpenHtmlCms)))
	http.HandleFunc("/profile", c.CookieUser(http.HandlerFunc(c.OpenHtmlProfile)))
	http.HandleFunc("/", c.OpenHtmlAddFavorites)
	http.HandleFunc("/404", c.OpenHtml404)
	http.HandleFunc("/create_user", c.OpenHtmlCreateUser)
	http.HandleFunc("/login_check", c.OpenHtmlLoginCheck)
	http.HandleFunc("/sales", c.OpenHtmlSales)
	http.HandleFunc("/prosv", c.OpenHtmlProsv)
	http.HandleFunc("/naura", c.OpenHtmlNaura)
	http.HandleFunc("/agat", c.OpenHtmlAgat)
	http.HandleFunc("/804", c.OpenHtml804)
	http.HandleFunc("/stronikum", c.OpenHtmlStronikum)

	// http.HandleFunc("/", c.SwitchUrl)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Starting server : [ %v ]\n", c.Configuration.Port)
	log.Fatal(http.ListenAndServe(":"+c.Configuration.Port, nil))
	// fmt.Printf("c.Configuration: %v\n", c.Configuration)
}

// func (c *Core) SwitchUrl(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.URL.Path)
// 	path := strings.Split(r.URL.Path, "/")

// 	switch r.URL.Path{
// 	case "/":
// 	case "/home":
// 	default:
// 		tmpl, _ := template.ParseFiles(rout.DataTemp.Path_prefix + rout.DataTemp.Path_frontend + "404.html")
// 		tmpl.Execute(w, rout.DataTemp)
// 	}
// 	if len(path) == 2 {
// 		if path[0] == "buy" {
// 			order_id := path[1]
// 			fmt.Println(order_id)
// 		}
// 	}
// 	fmt.Println(r.Cookies())
// }

func NewCore() *Core {
	c := config.NewConfiguration()
	// bd := bd.NewBd(*c)
	conn := connector.NewConnector("bookrzn", "book1995", "127.0.0.1:3306", "bookrzn")
	a := auth.NewAuth(*c, *conn)
	mw := middleware.NewMiddleware(*a)
	ps := parsing.NewParsingService(*c, *conn)

	fmt.Println(len(ps.ProsvCards))
	dt := datatemp.NewDataTemp(*c, ps.ProsvCards)
	rout := routes.NewRout(*a, *conn, *dt)

	return &Core{
		Configuration:  c,
		Connector:      conn,
		Auth:           a,
		Middleware:     mw,
		ParsingService: ps,
		Rout:           rout,
	}
}
