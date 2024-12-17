package handlers

import (
	"encoding/json"
	"net/http"
	"a21hc3NpZ25tZW50/service"
	"a21hc3NpZ25tZW50/utils"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user service.User // Pastikan struct User sudah terdefinisi dengan benar di dalam service
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Hash password sebelum disimpan
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Menyimpan user ke database
	err = service.CreateUser(user.Username, hashedPassword)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "user created successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user service.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Mengecek username dan password dari database
	storedUser, err := service.GetUserByUsername(user.Username)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Memeriksa apakah password cocok
	if !utils.CheckPasswords(storedUser.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Secret key untuk signing JWT
	secretKey := []byte("your_secret_key") // Ganti dengan key yang lebih aman

	// Generate JWT token
	token, err := utils.GenerateJWT(storedUser.Username, secretKey)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Mengembalikan token sebagai response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}