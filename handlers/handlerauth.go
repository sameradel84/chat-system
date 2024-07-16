package handlers

import (
	"chatsystem/models"
	"encoding/json"
	"net/http"

	"chatsystem/services"
)

// UserRegisterHandler handles HTTP requests to register a new user.
// It decodes the incoming JSON request body into a UserAuth model,
// calls the RegisterUserService to register the user,
// and responds with HTTP status 201 if successful.
func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserAuth
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.RegisterUserService(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// UserLoginHandler handles HTTP requests for user login.
// It decodes the incoming JSON request body into a UserAuth model,
// calls the LoginUserService to verify user credentials and generate a token,
// and responds with HTTP status 200 and the generated token if successful.
func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserAuth
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := services.LoginUserService(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
