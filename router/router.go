package router

import (
	"chatsystem/handlers"
	"log"
	"net/http"
)

func RouterStart() {
	http.HandleFunc("/register", handlers.UserRegisterHandler)
	http.HandleFunc("/login", handlers.UserLoginHandler)
	http.HandleFunc("/send", handlers.SendMessageHandler)
	http.HandleFunc("/messages", handlers.GetMessagesHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
