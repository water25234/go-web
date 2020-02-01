package middleware

import (
	http2 "ikdev/go-web/app/http"
	"ikdev/go-web/exception"
	"ikdev/go-web/helper"
	"net/http"
)

// Refresh JWT token
func (Middleware) RefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := http2.ServiceContainer.Invoke(func(a *helper.Auth) {
			a.RefreshToken()
		})

		if err != nil {
			exception.ProcessError(err)
		}

		next.ServeHTTP(w, r)
	})
}