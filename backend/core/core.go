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
	// http.HandleFunc("/favorites", c.CookieUser(http.HandlerFunc(c.OpenHtmlFavorites)))
	// http.HandleFunc("/orders", c.CookieUser(http.HandlerFunc(c.OpenHtmlOrders)))
	// http.HandleFunc("/fast_order", c.OpenHtmlFastOrder)
	// http.HandleFunc("/fast_order_save", c.OpenHtmlFastOrderSave)
	// http.HandleFunc("/home", c.OpenHtmlHome)
	// http.HandleFunc("/out", c.CookieUser(http.HandlerFunc(c.OpenHtmlout)))
	// http.HandleFunc("/login", c.OpenHtmlLogin)
	// http.HandleFunc("/registration", c.OpenHtmlRegistry)
	// http.HandleFunc("/cms", c.CookieAdmin(http.HandlerFunc(c.OpenHtmlCms)))
	// http.HandleFunc("/404", c.OpenHtml404)
	// http.HandleFunc("/create_user", c.OpenHtmlCreateUser)
	// http.HandleFunc("/login_check", c.OpenHtmlLoginCheck)
	// http.HandleFunc("/prosv", c.OpenHtmlProsv)
	// http.HandleFunc("/804", c.OpenHtml804)
	// http.HandleFunc("/search", c.OpenHtmlSearch)
	// http.HandleFunc("/for_school", c.OpenHtmlForSchool)

	http.HandleFunc("/", c.ServerRoutHtml) // c.CookieConfirm(http.HandlerFunc(c.ServerRoutHtml)))

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("[ ADDR ] -- [ %v ] -- [ %v ]", c.Configuration.Port, connector.DateTimeNow())
	fmt.Printf("\n[ SERVER ] -- [ START ] -- [ %v ]", connector.DateTimeNow())
	log.Fatal(http.ListenAndServe(":"+c.Configuration.Port, nil))

}

func NewCore() *Core {
	c := config.NewConfiguration()
	conn := connector.NewConnector(*c)
	a := auth.NewAuth(*c, *conn)
	mw := middleware.NewMiddleware(*a)
	ps := parsing.NewParsingService(*c, *conn)
	dt := datatemp.NewDataTemp(*c, ps.ListTargetCardCache)
	rout := routes.NewRout(*a, *c, *conn, *dt, *ps)

	return &Core{
		Configuration:  c,
		Connector:      conn,
		Auth:           a,
		Middleware:     mw,
		ParsingService: ps,
		Rout:           rout,
	}
}
