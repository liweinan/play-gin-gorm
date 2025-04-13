# Gin-GORM REST API

A RESTful API built with Gin and GORM, providing user management functionality with PostgreSQL database integration.

## Features

- User CRUD operations
- PostgreSQL database integration
- Environment-based configuration
- RESTful API endpoints

## Prerequisites

- Go 1.21 or higher
- PostgreSQL database
- Git

## Database Setup

### Option 1: Using PostgreSQL CLI

1. Connect to PostgreSQL:
   ```bash
   psql -U postgres
   ```

2. Create the database:
   ```sql
   CREATE DATABASE gin_gorm;
   ```

3. Create a user (if needed):
   ```sql
   CREATE USER your_username WITH PASSWORD 'your_password';
   ```

4. Grant privileges:
   ```sql
   GRANT ALL PRIVILEGES ON DATABASE gin_gorm TO your_username;
   ```

### Option 2: Using Docker

1. Pull the PostgreSQL image:
   ```bash
   docker pull postgres
   ```

2. Run the container:
   ```bash
   docker run --name gin-gorm-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=gin_gorm -p 5432:5432 -d postgres
   ```

### Option 3: Using pgAdmin (GUI)

1. Open pgAdmin and connect to your PostgreSQL server
2. Right-click on "Databases" and select "Create" > "Database"
3. Enter "gin_gorm" as the database name
4. Click "Save" to create the database

After creating the database, update your `.env` file with the correct credentials:
```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=gin_gorm
DB_PORT=5432
```

## Project Structure

```
.
├── config/         # Database configuration
├── handlers/       # Request handlers
├── models/         # Data models
├── .env           # Environment variables
├── go.mod         # Go module file
├── go.sum         # Go dependencies checksum
└── main.go        # Application entry point
```

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/gin-gorm.git
   cd gin-gorm
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Configure the database:
   - Create a PostgreSQL database
   - Copy the `.env.example` file to `.env`
   - Update the `.env` file with your database credentials:
     ```
     DB_HOST=localhost
     DB_USER=your_username
     DB_PASSWORD=your_password
     DB_NAME=your_database_name
     DB_PORT=5432
     ```

4. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## API Endpoints

### Users

- `POST /users` - Create a new user
- `GET /users` - Get all users
- `GET /users/:id` - Get a specific user
- `PUT /users/:id` - Update a user
- `DELETE /users/:id` - Delete a user

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [godotenv](https://github.com/joho/godotenv) - Environment variable loader

## License

This project is licensed under the MIT License - see the LICENSE file for details.
