package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/types"
	"github.com/asutosh29/amx-restro/pkg/utils/jwt_utils"
)

// For othar than login and register routes
func LoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for token
		if r.URL.Path == "/login" || r.URL.Path == "/register" || r.URL.Path == "/logout" {
			next.ServeHTTP(w, r)
		}
		cookie_jwt, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println("No Cookie found")
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jwt_token := cookie_jwt.Value
		claims, err := jwt_utils.ValidateJWT(jwt_token)
		if err != nil {
			fmt.Println("Error validating JWT")
			next.ServeHTTP(w, r)
		}

		// TODO: Store User in context for Frontend
		ctx := context.WithValue(r.Context(), "User", claims.User)
		// ctx = context.WithValue(ctx, "key", value)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// For login and register routes
func NewUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for token
		cookie_jwt, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println("No Cookie found")
				next.ServeHTTP(w, r)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jwt_token := cookie_jwt.Value
		_, err = jwt_utils.ValidateJWT(jwt_token)
		if err != nil {
			fmt.Println("Error validating JWT")
			next.ServeHTTP(w, r)
		}
		// TODO: Store User in context for Frontend
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	})
}

// AdminAccessOnly
func AdminAccessOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		User := r.Context().Value("User").(types.User)
		if !(User.Userole == "admin") && !(User.Userole == "super") {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		next.ServeHTTP(w, r)
	})
}
