package database

import (
	"chatsystem/models"
	"encoding/json"
	"log"
	"os"

	"github.com/gocql/gocql"
)

// CassandraConfig holds the configuration details required to connect to Cassandra.
type CassandraConfig struct {
	Hosts       []string `json:"hosts"`
	Keyspace    string   `json:"keyspace"`
	Consistency string   `json:"consistency"`
}

var Session *gocql.Session

// CassandraNewConn returns the active Cassandra session.
func CassandraNewConn() *gocql.Session {
	return Session
}

// InitCassandra initializes a connection to Cassandra using configuration from cassandra_config.json.
func InitCassandra() {
	configFile := "config/cassandra_config.json"
	configData, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Failed to read Cassandra config file: %v", err)
	}

	var config CassandraConfig
	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Failed to parse Cassandra config: %v", err)
	}

	cluster := gocql.NewCluster(config.Hosts...)
	cluster.Keyspace = config.Keyspace
	cluster.Consistency = gocql.Quorum
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatalf("Unable to connect to Cassandra: %v", err)
	}
	log.Println("Cassandra connected")
}

// FetchMessages retrieves messages from Cassandra for a given sender and recipient.
func FetchMessages(sender, recipient string) ([]models.Message, error) {
	var messages []models.Message
	iter := Session.Query(`SELECT sender, recipient, content, timestamp FROM messages WHERE sender = ? AND recipient = ?`, sender, recipient).Iter()
	var msg models.Message
	for iter.Scan(&msg.Sender, &msg.Recipient, &msg.Content, &msg.Timestamp) {
		messages = append(messages, msg)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return messages, nil
}
