package middleware

import (
	"backend/auth"
	"log"
	"net/http"
	"time"
)

type Middleware struct {
	auth.Auth
}

func NewMiddleware(a auth.Auth) *Middleware {
	return &Middleware{
		Auth: a,
	}
}

func (m *Middleware) CookieAdmin(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isFindCookie := m.GetCookieAdmin(w, r)
		if isFindCookie == true {
			next.ServeHTTP(w, r)
		}

	})
}

func (m *Middleware) CookieUser(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isFindCookie := m.GetCookieUser(w, r)
		if isFindCookie == true {
			next.ServeHTTP(w, r)
		}

	})
}

func (m *Middleware) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
