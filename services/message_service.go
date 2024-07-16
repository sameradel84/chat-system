package services

import (
	"chatsystem/database"
	"chatsystem/models"
	"encoding/json"
	"fmt"
)

// SendMessageService inserts a message into the database and invalidates the cache for the sender and recipient.
func SendMessageService(msg *models.MessageUser) error {
	// Insert message into Cassandra database
	if err := database.Session.Query(`INSERT INTO messages (sender, recipient, content, timestamp) VALUES (?, ?, ?, toTimestamp(now()))`, msg.Sender, msg.Recipient, msg.Content).Exec(); err != nil {
		return err
	}

	// Invalidate cache for messages between sender and recipient
	cacheKey := fmt.Sprintf("messages:%s:%s", msg.Sender, msg.Recipient)
	database.InvalidateCache(cacheKey)
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
