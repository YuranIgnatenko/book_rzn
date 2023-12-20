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
	// c.CookieConfirm(http.HandlerFunc(c.ServerRoutHtml))

	http.HandleFunc("/", c.ServerSwitchRout)

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
