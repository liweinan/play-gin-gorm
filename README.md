# Gin-GORM REST API

A RESTful API built with Gin and GORM, providing user management functionality with PostgreSQL database integration.

## Features

- User CRUD operations
- PostgreSQL database integration
- Docker containerization
- RESTful API endpoints
- Health checks for database availability

## Prerequisites

- Docker and Docker Compose
- Git

## Project Structure

```
.
├── config/         # Database configuration
├── handlers/       # Request handlers
├── models/         # Data models
├── .env           # Environment variables
├── Dockerfile     # Application container configuration
├── docker-compose.yml # Docker services configuration
├── go.mod         # Go module file
├── go.sum         # Go dependencies checksum
└── main.go        # Application entry point
```

## Quick Start with Docker

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/gin-gorm.git
   cd gin-gorm
   ```

2. Start the application using Docker Compose:
   ```bash
   docker-compose up -d
   ```

The application will be available at `http://localhost:8080`

## Docker Configuration

The project uses Docker Compose to run two containers:

1. **PostgreSQL Database (gin-gorm-db)**
   - Port: 5432
   - Username: postgres
   - Password: postgres
   - Database: gin_gorm
   - Includes health checks to ensure database availability

2. **Go Application (gin-gorm-app)**
   - Port: 8080
   - Waits for database to be healthy before starting

## Docker Commands

- Start the application:
  ```bash
  docker-compose up -d
  ```

- Stop the application:
  ```bash
  docker-compose down
  ```

- View logs:
  ```bash
  docker-compose logs -f
  ```

- Rebuild and restart:
  ```bash
  docker-compose up -d --build
  ```

- Recreate containers (useful when you need to reset the database):
  ```bash
  # Stop and remove all containers
  docker-compose down

  # Remove the volume to delete all data
  docker volume rm play-gin-gorm_postgres_data

  # Start fresh containers
  docker-compose up -d --build
  ```

  Note: This will delete all data in the database. Use this command when you want to start with a clean database.

## API Endpoints

### Users

- `POST /users/` - Create a new user
  ```bash
  # Required fields:
  # - username (must be unique)
  # - email (must be unique)
  # - password
  curl -X POST -H "Content-Type: application/json" \
    -d '{"username": "testuser", "email": "test@example.com", "password": "your_password"}' \
    http://localhost:8080/users/
  ```

  Example response:
  ```json
  {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "created_at": "2025-04-13T12:03:51.880717168Z",
    "updated_at": "2025-04-13T12:03:51.880717168Z"
  }
  ```

- `GET /users/` - Get all users
  ```bash
  curl http://localhost:8080/users/
  ```

- `GET /users/:id` - Get a specific user
  ```bash
  curl http://localhost:8080/users/1
  ```

- `PUT /users/:id` - Update a user
  ```bash
  curl -X PUT -H "Content-Type: application/json" \
    -d '{"name": "Updated User", "email": "updated@example.com"}' \
    http://localhost:8080/users/1
  ```

- `DELETE /users/:id` - Delete a user
  ```bash
  curl -X DELETE http://localhost:8080/users/1
  ```

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) v1.9.1 - Web framework
- [GORM](https://gorm.io/) v1.25.5 - ORM library
- [PostgreSQL](https://www.postgresql.org/) - Database
- [godotenv](https://github.com/joho/godotenv) v1.5.1 - Environment variable loader

## Development

For local development without Docker:

1. Install PostgreSQL locally
2. Copy `.env.example` to `.env` and update the credentials
3. Install Go dependencies:
   ```bash
   go mod download
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
