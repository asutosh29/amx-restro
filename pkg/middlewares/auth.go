package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/asutosh29/amx-restro/pkg/types"
	"github.com/asutosh29/amx-restro/pkg/utils/jwt_utils"
)

// For login and register routes
func NewUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for token
		cookie_jwt, err := r.Cookie("token")
		if err == nil {
			fmt.Println("Cookie found")

			jwt_token := cookie_jwt.Value
			jwt_claims, err := jwt_utils.ValidateJWT(jwt_token)
			if err == nil {
				fmt.Println("Valid JWT")
				if IsValidUser, _ := models.UserExistsById(jwt_claims.User.UserId); IsValidUser {
					fmt.Println("Valid user")
					http.Redirect(w, r, "/home", http.StatusSeeOther)
					return
				} else {
					next.ServeHTTP(w, r)
				}
			}

			next.ServeHTTP(w, r)
		}

		next.ServeHTTP(w, r)
	})
}

// AdminAccessOnly
func AdminAccessOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		User := r.Context().Value("User").(types.User)
		if !(User.Userole == types.ROLE().ADMIN) && !(User.Userole == types.ROLE().SUPER) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		next.ServeHTTP(w, r)
	})
}

func RestrictToLoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// No Cookie redirect to Log in
		cookie_jwt, err := r.Cookie("token")
		if err != nil {
			fmt.Println("No Cookie found")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		// Invalid JWT redirect to Log in
		jwt_token := cookie_jwt.Value
		jwt_claims, err := jwt_utils.ValidateJWT(jwt_token)
		if err != nil {
			fmt.Println("Error validating JWT")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// If not user then Login

		if IsValidUser, _ := models.UserExistsById(jwt_claims.User.UserId); !IsValidUser {
			fmt.Println("Invalid user")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		// fmt.Println("Claims: ", jwt_claims.User.UserId)
		ctx := context.WithValue(r.Context(), "User", jwt_claims.User)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
