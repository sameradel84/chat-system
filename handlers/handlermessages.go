package handlers

import (
	"chatsystem/models"
	"chatsystem/services"
	"encoding/json"
	"net/http"
)

// SendMessageHandler handles HTTP requests to send a message.
// It decodes the incoming JSON request body into a MessageUser model,
// calls the SendMessageService to process the message,
// and responds with HTTP status 201 if successful.
func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var msg models.MessageUser
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.SendMessageService(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetMessagesHandler handles HTTP requests to retrieve messages between two users.
// It retrieves sender and recipient from query parameters,
// calls the GetMessagesService to fetch messages,
// and responds with HTTP status 200 and JSON array of messages if successful.
func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	sender := r.URL.Query().Get("sender")
	recipient := r.URL.Query().Get("recipient")

	messages, err := services.GetMessagesService(sender, recipient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
