package middleware

import (
	"backend/auth"
	"fmt"
	"log"
	"net/http"
	"time"
)

func CookieAdmin(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware -- ok A admin")
		isFindCookie := auth.GetCookieAdmin(w, r)
		if isFindCookie == true {
			next.ServeHTTP(w, r)
		}

		fmt.Println("middleware -- ok B admin")
	})
}

func CookieUser(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware -- ok A")
		isFindCookie := auth.GetCookieUser(w, r)
		if isFindCookie == true {
			next.ServeHTTP(w, r)
		}

		fmt.Println("middleware -- ok B")
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
