package services

import (
	"chatsystem/database"
	"chatsystem/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

// SendMessageService inserts a message into the database and invalidates the cache for the sender and recipient.
func SendMessageService(msg *models.MessageUser) error {
	// Insert message into Cassandra database
	query := `INSERT INTO chatsystem.messages (id, sender, recipient, content, timestamp) VALUES (?, ?, ?, ?, ?)`
	if err := database.Session.Query(query, gocql.TimeUUID(), msg.Sender, msg.Recipient, msg.Content, time.Now().UTC()).Exec(); err != nil {
		return err
	}

	// Invalidate cache for messages between sender and recipient
	cacheKeySender := fmt.Sprintf("messages:%s", msg.Sender)
	cacheKeyRecipient := fmt.Sprintf("messages:%s", msg.Recipient)
	if err := database.InvalidateCache(cacheKeySender, cacheKeyRecipient); err != nil {
		return err
	}

	return nil
}

// GetMessagesService retrieves messages between sender and recipient.
// It first tries to fetch messages from cache; if not found, retrieves from the database and updates the cache.
func GetMessagesService(sender, recipient string) ([]models.Message, error) {
	// Generate cache key for messages between sender and recipient
	cacheKey := fmt.Sprintf("messages:%s:%s", sender, recipient)

	// Try to fetch messages from cache
	cachedMessages, err := database.GetCache(cacheKey)
	if err == nil && cachedMessages != "" {
		// Unmarshal cached messages into []models.Message
		var messages []models.Message
		if err := json.Unmarshal([]byte(cachedMessages), &messages); err == nil {
			return messages, nil
		}
	}

	// Fetch messages from database if not found in cache
	messages, err := database.FetchMessages(sender, recipient)
	if err != nil {
		return nil, err
	}

	// Marshal messages to JSON bytes and set cache
	messagesBytes, _ := json.Marshal(messages)
	database.SetCache(cacheKey, string(messagesBytes))
	return messages, nil
}
