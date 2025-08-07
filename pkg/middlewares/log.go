package middlewares

import (
	"fmt"
	"net/http"
)

func LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v : %v\n", r.URL, r.Method)

		next.ServeHTTP(w, r)
	})
}
