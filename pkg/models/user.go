package models

import (
	"fmt"

	"github.com/asutosh29/amx-restro/pkg/types"
)

func IsFirstUser(user types.User) bool {
	row := DB.QueryRow("select email, username from users")
	var email string
	var username string
	row.Scan(&email, &username)
	if email == "" || username == "" {
		return true
	}
	return false
}

func IsEmailUnique(user types.User) (bool, error) {
	var email string
	err := DB.QueryRow("select email from users where email=?", user.Email).Scan(&email)
	if err != nil {
		return true, err
	}
	if email == "" {
		return true, nil
	}
	return false, nil
}

func IsUserNameUnique(user types.User) (bool, error) {
	var username string
	err := DB.QueryRow("select username from users where username=?", user.Username).Scan(&username)
	if err != nil {
		return true, err
	}
	if username == "" {
		return true, nil
	}
	return false, nil
}

func AddUser(user types.User) error {
	_, err := DB.Exec("insert into users(email, username,userRole,first_name,last_name,contact, hashpwd) values(?,?,?,?,?,?,?)", user.Email, user.Username, user.Userole, user.First_name, user.Last_name, user.Contact, user.Hashpwd)
	if err != nil {
		return err
	}
	fmt.Println("User added Successfully")
	return nil
}

func GetUser(user types.User) (types.User, error) {
	var DbUser types.User
	err := DB.QueryRow("select email, username, first_name, last_name, contact, hashpwd, userRole from users where email=?", user.Email).Scan(&DbUser.Email, &DbUser.Username, &DbUser.First_name, &DbUser.Last_name, &DbUser.Contact, &DbUser.Hashpwd, &DbUser.Userole)
	if err != nil {
		return DbUser, err
	}
	return DbUser, nil
}
