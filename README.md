# User CRUD REST API with Go, Mux, and PostgreSQL

This is a simple RESTful API for managing users (name and address) using Go, the `gorilla/mux` router, and PostgreSQL as the database.

## Features

- Create a user
- Get a user by ID
- Update a user
- Delete a user
- List all users

## Tech Stack

- **Go**: Programming language
- **gorilla/mux**: HTTP router and URL matcher for building Go web applications
- **PostgreSQL**: Relational database for storing user information
- **Docker** (optional): Containerize the application and the database

## Project Structure

```bash
.
├── main.go                # Entry point of the application
├── controller
│   └── user_controller.go # Contains handler functions for user routes
├── model
│   └── user.go            # Defines the User struct and database operations
├── config
│   └── db.go              # PostgreSQL database connection setup
├── router
│   └── router.go          # Defines all the application routes
├── go.mod                 # Go module dependencies
└── README.md              # Project documentation
```

## Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/YOUR_USERNAME/YOUR_REPOSITORY.git
cd YOUR_REPOSITORY
```

### 2. Install Dependencies

Make sure you have Go installed, then run:

```bash
go mod tidy
```

This will download the necessary dependencies (`gorilla/mux` and `lib/pq` for PostgreSQL).

### 3. Set Up PostgreSQL Database

Create a PostgreSQL database and a table for users.

```sql
CREATE DATABASE userdb;

\c userdb

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    address TEXT
);
```

### 4. Set Up Environment Variables

Create a `.env` file in the root directory and add your PostgreSQL configuration:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=userdb
```

### 5. Run the Application

```bash
go run main.go
```

The API will start on `localhost:8000`.

### 6. API Endpoints

| Method | Endpoint           | Description           | Body (JSON)             |
|--------|--------------------|-----------------------|-------------------------|
| GET    | `/users`            | Get all users         |                         |
| GET    | `/users/{id}`       | Get user by ID        |                         |
| POST   | `/users`            | Create a new user     | `{ "name": "John", "address": "123 St" }` |
| PUT    | `/users/{id}`       | Update a user         | `{ "name": "John", "address": "123 St" }` |
| DELETE | `/users/{id}`       | Delete a user         |                         |

### 7. Docker (Optional)

To run the application and PostgreSQL using Docker, follow these steps:

1. Build the Docker image:
   ```bash
   docker build -t user-crud-api .
   ```

2. Run the Docker container with PostgreSQL and the application:
   ```bash
   docker-compose up
   ```

This will spin up both the application and PostgreSQL in Docker containers.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
