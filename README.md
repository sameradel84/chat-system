# Chatsystem

Chatsystem is a microservice-based chat application built with Go and integrated with Cassandra for message storage and Redis for caching.

## Features

- User registration and authentication
- Sending messages between users
- Retrieving message history
- Caching messages for efficient retrieval

## Technologies Used

- Golang
- Cassandra
- Redis
- Docker
- NGINX

## Installation and Setup

### Prerequisites

- Docker
- Docker Compose
- go

### Running the Application

1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd chatsystem

2. Start the application using Docker Compose:
`docker-compose up`

3. The application will be accessible at http://localhost:8080.


## API Endpoints
### User Registration
- **Endpoint:** /register
- **Method:** POST
- **Request Body:**
``` json
{
  "username": "user1",
  "password": "password123"
}
```

