package main

import (
	"chatsystem/database"

	"chatsystem/router"
)

func main() {
	// Initialize Redis
	database.InitRedis()

	// Initialize Cassandra
	database.InitCassandra()

	// Initialize server connection and endpoints
	router.RouterStart()

}
