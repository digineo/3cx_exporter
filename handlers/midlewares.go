package handlers

import (
	"net/http"
)

func getRequestCountMidleware(provisor stateProvisor) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			provisor.RegisterRequest()
			next.ServeHTTP(w, r)
		})
	}

}
