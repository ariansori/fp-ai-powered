package service

import (
	"database/sql"
	"a21hc3NpZ25tZW50/database"
	"fmt"
)

// Struct untuk User
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Fungsi untuk menambahkan user baru
func CreateUser(username, password string) error {
	// Menyimpan user ke dalam database
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := database.DB.Exec(query, username, password)
	if err != nil {
		return fmt.Errorf("unable to create user: %v", err)
	}
	return nil
}

// Fungsi untuk mendapatkan user berdasarkan username
func GetUserByUsername(username string) (User, error) {
	var user User
	query := "SELECT id, username, password FROM users WHERE username = $1"
	row := database.DB.QueryRow(query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("user not found")
		}
		return User{}, fmt.Errorf("unable to retrieve user: %v", err)
	}
	return user, nil
}