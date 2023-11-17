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
	http.HandleFunc("/home", c.CookieUser(http.HandlerFunc(c.OpenHtmlHome)))
	http.HandleFunc("/logout", c.CookieUser(http.HandlerFunc(c.OpenHtmlLogout)))
	http.HandleFunc("/new", c.CookieUser(http.HandlerFunc(c.OpenHtmlNew)))
	http.HandleFunc("/payment", c.CookieUser(http.HandlerFunc(c.OpenHtmlPayment)))
	http.HandleFunc("/collectschool", c.CookieUser(http.HandlerFunc(c.OpenHtmlCollectschool)))
	http.HandleFunc("/login", c.OpenHtmlLogin)
	http.HandleFunc("/registration", c.OpenHtmlRegistry)
	http.HandleFunc("/cms", c.CookieAdmin(http.HandlerFunc(c.OpenHtmlCms)))
	http.HandleFunc("/404", c.OpenHtml404)
	http.HandleFunc("/create_user", c.OpenHtmlCreateUser)
	http.HandleFunc("/login_check", c.OpenHtmlLoginCheck)
	http.HandleFunc("/sales", c.OpenHtmlSales)
	http.HandleFunc("/prosv", c.OpenHtmlProsv)
	http.HandleFunc("/naura", c.OpenHtmlNaura)
	http.HandleFunc("/agat", c.OpenHtmlAgat)
	http.HandleFunc("/804", c.OpenHtml804)
	http.HandleFunc("/stronikum", c.OpenHtmlStronikum)
	http.HandleFunc("/search", c.OpenHtmlSearch)

	// hanlder: favorites, profiles, orders, targets ..
	// unique content for users
	http.HandleFunc("/", c.CookieUser(http.HandlerFunc(c.OpenHtmlProfile)))

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Starting server : [ %v ]\n", c.Configuration.Port)
	log.Fatal(http.ListenAndServe(":"+c.Configuration.Port, nil))

}

func NewCore() *Core {
	c := config.NewConfiguration()
	conn := connector.NewConnector(*c)
	a := auth.NewAuth(*c, *conn)
	mw := middleware.NewMiddleware(*a)
	ps := parsing.NewParsingService(*c, *conn)

	dt := datatemp.NewDataTemp(*c, ps.TargetCards)
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
