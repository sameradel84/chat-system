package services

import (
	"chatsystem/database"
	"chatsystem/models"
	"chatsystem/utils"
	"errors"
)

// RegisterUserService registers a new user by validating credentials,
// hashing the password, and storing the user in the database.
func RegisterUserService(user *models.UserAuth) error {
	// Validate user credentials (username and password)
	if err := utils.ValidateUserCredentials(user.Username, user.Password); err != nil {
		return err
	}

	// Hash the user's password before storing it in the database
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	// Insert user into the database (users table)
	if err := database.Session.Query(`INSERT INTO users (username, password) VALUES (?, ?)`, user.Username, user.Password).Exec(); err != nil {
		return err
	}

	return nil
}

// LoginUserService verifies user credentials by validating username and password,
// fetching the stored user from the database, and checking the password hash.
// If credentials are valid, it generates and returns an authentication token.
func LoginUserService(user *models.UserAuth) (string, error) {
	// Validate user credentials (username and password)
	if err := utils.ValidateUserCredentials(user.Username, user.Password); err != nil {
		return "", err
	}

	// Retrieve stored user from the database
	var storedUser models.UserAuth
	if err := database.Session.Query(`SELECT username, password FROM users WHERE username = ? LIMIT 1`, user.Username).Scan(&storedUser.Username, &storedUser.Password); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Check if the provided password matches the stored password hash
	if !utils.CheckPasswordHash(user.Password, storedUser.Password) {
		return "", errors.New("invalid username or password")
	}

	// Generate a dummy token
	token := "dummy-token"
	return token, nil
}
