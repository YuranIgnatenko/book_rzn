package main

import (
	"backend/config"
	"backend/middleware"
	"backend/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("init server")
	Conf := config.NewConfiguration()

	http.HandleFunc("/about", routes.OpenHtmlAbout)
	http.HandleFunc("/blog", routes.OpenHtmlBlog)
	http.HandleFunc("/cart", middleware.CookieUser(http.HandlerFunc(routes.OpenHtmlCart)))
	http.HandleFunc("/contacts", routes.OpenHtmlContacts)
	http.HandleFunc("/delivery", middleware.CookieUser(http.HandlerFunc(routes.OpenHtmlDelivery)))
	http.HandleFunc("/exchange", middleware.CookieUser(http.HandlerFunc(routes.OpenHtmlExchange)))
	http.HandleFunc("/favorites", middleware.CookieUser(http.HandlerFunc(routes.OpenHtmlFavorites)))
	http.HandleFunc("/home", routes.OpenHtmlHome)
	http.HandleFunc("/new", middleware.CookieUser(http.HandlerFunc(routes.OpenHtmlNew)))
	http.HandleFunc("/payment", middleware.CookieUser(http.HandlerFunc(routes.OpenHtmlPayment)))
	http.HandleFunc("/collectschool", middleware.CookieUser(http.HandlerFunc(routes.OpenHtmlCollectschool)))
	http.HandleFunc("/login", routes.OpenHtmlLogin)
	http.HandleFunc("/registration", routes.OpenHtmlRegistry)
	http.HandleFunc("/cms", middleware.CookieAdmin(http.HandlerFunc(routes.OpenHtmlCms)))
	http.HandleFunc("/profile", middleware.CookieUser(http.HandlerFunc(routes.OpenHtmlProfile)))
	http.HandleFunc("/buy_order", routes.OpenHtmlBuyOrder)
	http.HandleFunc("/404", routes.OpenHtml404)
	http.HandleFunc("/create_user", routes.OpenHtmlCreateUser)
	http.HandleFunc("/login_check", routes.OpenHtmlLoginCheck)
	http.HandleFunc("/sales", routes.OpenHtmlSales)
	http.HandleFunc("/prosv", routes.OpenHtmlProsv)
	http.HandleFunc("/naura", routes.OpenHtmlNaura)
	http.HandleFunc("/agat", routes.OpenHtmlAgat)
	http.HandleFunc("/804", routes.OpenHtml804)
	http.HandleFunc("/stronikum", routes.OpenHtmlStronikum)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Starting server : [ %v ]\n", Conf.Port)
	log.Fatal(http.ListenAndServe(":"+Conf.Port, nil))

}
