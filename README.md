# Chat system

## Overview

Chat System is a microservices-based chat application built using Golang, Cassandra, Redis, and Docker. This architecture ensures high performance, scalability, and maintainability.

## Architectural Decisions

1. **Microservices Architecture**:
   - Ensures modularity, scalability, and maintainability.
   - Independent development, deployment, and scaling of services.

2. **Golang for Backend**:
   - Chosen for performance, concurrency model, and strong static typing.
   - Suitable for building high-performance, concurrent server-side applications.

3. **Cassandra for Data Storage**:
   - Used as the primary database for its high availability, scalability, and minimal latency.
   - Handles large volumes of data efficiently.

4. **Redis for Caching**:
   - Caches frequently accessed data to reduce load on Cassandra and improve response times.
   - Used for caching user sessions and ephemeral data.

5. **Docker for Containerization**:
   - Ensures consistency across different environments (development, testing, production).
   - Simplifies deployment and scaling.

6. **nginx as a Reverse Proxy**:
   - Handles incoming HTTP requests and routes them to the appropriate services.
   - Provides load balancing and can serve static content if necessary.

7. **Docker Compose for Orchestration**:
   - Defines and manages multi-container Docker applications.
   - Facilitates setup of development and testing environments.

## Assumptions

1. **Single Data Center Deployment**:
   - Initial deployment is within a single data center.
   - Can be extended to multiple data centers if needed.

2. **Simple Authentication**:
   - Basic authentication implemented for simplicity.
   - Can be replaced with more robust mechanisms (e.g., OAuth) in the future.

3. **Linear Scalability Needs**:
   - Assumes linear scalability needs.
   - Can scale horizontally by adding more service instances.

4. **High Availability**:
   - High availability is crucial, hence the use of Cassandra and Redis.

5. **Network Reliability**:
   - Assumes reliable underlying network.
   - Network issues handled at Docker and Kubernetes level.

## Services

- **app**: Main Golang application handling business logic.
- **cassandra**: Primary database for storing chat messages and user data.
- **redis**: Caching layer for improving performance and session management.
- **nginx**: Reverse proxy for routing HTTP requests and load balancing.

## Getting Started

### Technologies Used

- Golang
- Cassandra
- Redis
- Docker
- NGINX

### Installation and Setup

#### Prerequisites

- Docker
- Docker Compose
- go

### Running the Application

1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd chatsystem

2. Start the application using Docker Compose:
`docker-compose up --build`

3. The application will be accessible at http://localhost:8080.

4. The NGINX will be accessible at http://localhost:8081.

5. To stop and remove all containers:

    `docker-compose down`


## API Endpoints
### User Registration
- **Endpoint:** http://localhost:8080/register
- **Method:** POST
- **Request Body:**
``` json
{
  "username": "samer",
  "password": "password123"
}
```
### Response:

```json

{
    "message": "User registered successfully"
}
```
### User Login
- **URL:** http://localhost:8080/login
- **Method:** POST
- **Request Body:**

```json

{
    "username": "samer",
    "password": "password123"
}
```
### Response:

```json

{
    "token": "dummy-token"
}
```
### Send Message
- **URL:** http://localhost:8080/send
- **Method:** POST
- **Request Body:**

```json

{
    "sender": "sender_username",
    "recipient": "recipient_username",
    "content": "your_message"
}
```
### Response: 201 Created

### Get Messages
- **URL:** http://localhost:8080/messages
- **Method:** GET
- **Query Parameters: sender=sender_username&recipient=recipient_username**
### Response:

```json

[
    {
        "sender": "sender_username",
        "recipient": "recipient_username",
        "content": "your_message",
        "timestamp": "message_timestamp"
    }
]
```

## Author
This project was developed by [samer adel]

