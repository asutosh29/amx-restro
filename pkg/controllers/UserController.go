package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/asutosh29/amx-restro/pkg/types"
	"github.com/asutosh29/amx-restro/pkg/utils/hashing"
	"github.com/asutosh29/amx-restro/pkg/utils/jwt_utils"
)

func HandleRegisterUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	hashpwd := hashing.GenerateHashFromPassword(r.FormValue("password"))
	user := types.User{
		Username:  r.FormValue("username"),
		Email:     r.FormValue("email"),
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Contact:   r.FormValue("contact"),
		Hashpwd:   hashpwd,
		Userole:   types.ROLE().CUSTOMER,
	}

	if user.Username == "" || user.Email == "" || user.FirstName == "" || user.LastName == "" || user.Contact == "" || user.Hashpwd == "" {
		fmt.Println("Bad user input")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		// TODO: How to pass message to the front end?
		// NOTE: Store in local storage and render it on Front end from there! or try cookies.
		return
	}
	// Hash the Password

	// TODO: Implement first user is Super user
	if models.IsFirstUser(user) {
		user.Userole = types.ROLE().SUPER
		fmt.Println("Super User waas added")
		models.AddUser(user)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// TODO: Implement with DB action
	if val, err := models.IsEmailUnique(user); err != nil {
		fmt.Println("Email check error")
		fmt.Println(err)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	} else {
		if !val {
			fmt.Println("Email Already exists!")
			// TODO: Flash messages for frontend
			RenderRegister(w, r)
			return
		}
	}
	if val, err := models.IsUserNameUnique(user); err != nil {
		fmt.Println("Username check error")
		fmt.Println(err)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	} else {
		if !val {
			fmt.Println("Username Already exists!")
			// TODO: Flash messages for frontend
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}
	}

	// Add user to the DB
	err := models.AddUser(user)
	if err != nil {
		fmt.Println("Error adding user to Database")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}
	// TODO: Flash - User registered successfully
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func HandleLoginUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	user := types.User{
		Email: r.FormValue("email"),
	}

	if user.Email == "" || r.FormValue("password") == "" {
		fmt.Println("Bad user input")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		// TODO: How to pass message to the front end?
		// NOTE: Store in local storage and render it on Front end from there! or try cookies.
		return
	}

	if val, err := models.IsEmailUnique(user); err != nil {
		fmt.Println("Email check error")
		fmt.Println(err)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	} else {
		if val {
			fmt.Println("Email Already exists!")
			// TODO: Flash messages for frontend
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}

	}

	// Get Real User
	var RealUser types.User
	RealUser, err := models.GetUser(user)
	if err != nil {
		fmt.Println("Error retrieving user")
		fmt.Println(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Check Password
	if validPassword := hashing.CheckPasswordFromHash(RealUser.Hashpwd, r.FormValue("password")); !validPassword {
		// TODO: Flash message - Invalid Password
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Check JWT
	token_JWT, err := jwt_utils.GenerateJWT(RealUser)
	if err != nil {
		fmt.Println("Error Generating JWT token")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Send JWT in cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token_JWT,
		Expires: time.Now().Add(24 * time.Hour),
	})

	// Redirect to Home page
	http.Redirect(w, r, "/home", http.StatusSeeOther)

}

func HandleLogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  " ",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
