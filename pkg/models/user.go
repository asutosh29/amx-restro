package models

import (
	"fmt"

	"github.com/asutosh29/amx-restro/pkg/types"
)

func IsFirstUser(user types.User) bool {
	row := DB.QueryRow(`
    SELECT email, username
    FROM users
`)
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
	err := DB.QueryRow(`
    SELECT email
    FROM users
    WHERE email = ?
`, user.Email).Scan(&email)
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
	err := DB.QueryRow(`
    SELECT username
    FROM users
    WHERE username = ?
`, user.Username).Scan(&username)
	if err != nil {
		return true, err
	}
	if username == "" {
		return true, nil
	}
	return false, nil
}

func AddUser(user types.User) error {
	_, err := DB.Exec(`
    INSERT INTO users (email, username, userRole, first_name, last_name, contact, hashpwd)
    VALUES (?, ?, ?, ?, ?, ?, ?)
`, user.Email, user.Username, user.Userole, user.FirstName, user.LastName, user.Contact, user.Hashpwd)
	if err != nil {
		return err
	}
	fmt.Println("User added Successfully")
	return nil
}

func GetAllUsers() ([]types.User, error) {
	var AllUsers []types.User
	rows, err := DB.Query(`
    SELECT id, email, username, first_name, last_name, contact, hashpwd, userRole
    FROM users
`)
	if err != nil {
		return AllUsers, err
	}
	defer rows.Close()
	for rows.Next() {
		var DbUser types.User
		rows.Scan(&DbUser.UserId, &DbUser.Email, &DbUser.Username, &DbUser.FirstName, &DbUser.LastName, &DbUser.Contact, &DbUser.Hashpwd, &DbUser.Userole)
		AllUsers = append(AllUsers, DbUser)
	}
	return AllUsers, nil
}

func GetUser(user types.User) (types.User, error) {
	var DbUser types.User
	err := DB.QueryRow(`
    SELECT id, email, username, first_name, last_name, contact, hashpwd, userRole
    FROM users
    WHERE email = ?
`, user.Email).Scan(&DbUser.UserId, &DbUser.Email, &DbUser.Username, &DbUser.FirstName, &DbUser.LastName, &DbUser.Contact, &DbUser.Hashpwd, &DbUser.Userole)
	if err != nil {
		return DbUser, err
	}
	return DbUser, nil
}

func MakeAdminById(userId int) (int, error) {
	_, err := DB.Exec(`
    UPDATE users
    SET userRole = ?
    WHERE id = ?
`, types.ROLE().ADMIN, userId)
	if err != nil {
		fmt.Println("Error Making User as Admin")
		return userId, err
	}
	return userId, nil
}

func MakeCustomerById(userId int) (int, error) {
	_, err := DB.Exec(`
    UPDATE users
    SET userRole = ?
    WHERE id = ?
`, types.ROLE().CUSTOMER, userId)
	if err != nil {
		fmt.Println("Error Making User as Admin")
		return userId, err
	}
	return userId, nil
}
