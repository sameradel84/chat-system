package utils

import "errors"

// ValidateUserCredentials checks if the username and password are valid.
func ValidateUserCredentials(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username or password cannot be empty")
	}
	return nil
}
