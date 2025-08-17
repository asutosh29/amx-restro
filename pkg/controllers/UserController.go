package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/asutosh29/amx-restro/pkg/types"
	"github.com/asutosh29/amx-restro/pkg/utils/hashing"
	"github.com/asutosh29/amx-restro/pkg/utils/jwt_utils"
	"github.com/asutosh29/amx-restro/pkg/utils/session_utils"
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
		session_utils.FlashMsgErr(w, r, "Bad User Input!", true)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		// NOTE: Store in local storage and render it on Front end from there! or try cookies.
		return
	}

	if models.IsFirstUser(user) {
		user.Userole = types.ROLE().SUPER
		fmt.Println("Super User waas added")
		session_utils.FlashMsgErr(w, r, "Super User created!", false)
		models.AddUser(user)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if val, err := models.IsEmailUnique(user); err != nil {
		fmt.Println("New user spotted")
		fmt.Println(err)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	} else {
		if !val {
			fmt.Println("Email already registered")
			session_utils.FlashMsgErr(w, r, "Email Already Exists! Try another email", true)
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}
	}

	if val, err := models.IsUserNameUnique(user); err != nil {
		fmt.Println("New Username spotted")
	} else {
		if !val {
			fmt.Println("Username Already exists!")
			session_utils.FlashMsgErr(w, r, "Username taken. Try another username!", true)
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}
	}

	err := models.AddUser(user)
	if err != nil {
		fmt.Println("Error adding user to Database: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		session_utils.FlashMsgErr(w, r, "Error registering user", true)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	session_utils.FlashMsgErr(w, r, "User Registered Successfully", false)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func HandleLoginUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	user := types.User{
		Email: r.FormValue("email"),
	}

	if user.Email == "" || r.FormValue("password") == "" {
		fmt.Println("Bad user input")
		session_utils.FlashMsgErr(w, r, "Bad user input", true)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Unique TRUE => User doesn't exist
	if val, err := models.IsEmailUnique(user); err != nil {
		fmt.Println("Email check error")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		session_utils.FlashMsgErr(w, r, "No User found register first!", true)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	} else {
		if val {
			fmt.Println("Email doesn't exists!")
			session_utils.FlashMsgErr(w, r, "Email ID doesn't exist please register first", true)
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}

	}

	var RealUser types.User
	RealUser, err := models.GetUser(user)
	if err != nil {
		fmt.Println("Error retrieving user")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		session_utils.FlashMsgErr(w, r, "Error Retrieving User. Please Login with Valid details", true)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Check Password
	if validPassword := hashing.CheckPasswordFromHash(RealUser.Hashpwd, r.FormValue("password")); !validPassword {
		session_utils.FlashMsgErr(w, r, "Invalid Password. Try again", true)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Check JWT
	token_JWT, err := jwt_utils.GenerateJWT(RealUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		session_utils.FlashMsgErr(w, r, "Error Generating JWT", true)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Send JWT in cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token_JWT,
		Expires: time.Now().Add(24 * time.Hour),
	})

	session_utils.FlashMsgErr(w, r, "User logged in successfully", false)
	http.Redirect(w, r, "/home", http.StatusSeeOther)

}

func HandleLogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  " ",
		MaxAge: -1,
	})
	session_utils.FlashMsgErr(w, r, "Logged out successfully", false)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
