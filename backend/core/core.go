package core

import (
	"backend/config"
	"backend/middleware"
	"backend/routes"
	"fmt"
	"log"
	"net/http"
)

type Core struct {
	routes.Router
	middleware.Middleware
	config.Configuration
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
	http.HandleFunc("/buy_order", c.OpenHtmlBuyOrder)
	http.HandleFunc("/404", c.OpenHtml404)
	http.HandleFunc("/create_user", c.OpenHtmlCreateUser)
	http.HandleFunc("/login_check", c.OpenHtmlLoginCheck)
	http.HandleFunc("/sales", c.OpenHtmlSales)
	http.HandleFunc("/prosv", c.OpenHtmlProsv)
	http.HandleFunc("/naura", c.OpenHtmlNaura)
	http.HandleFunc("/agat", c.OpenHtmlAgat)
	http.HandleFunc("/804", c.OpenHtml804)
	http.HandleFunc("/stronikum", c.OpenHtmlStronikum)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Starting server : [ %v ]\n", c.Port)
	log.Fatal(http.ListenAndServe(":"+c.Port, nil))

}

func NewCore() *Core {
	return &Core{}
}
