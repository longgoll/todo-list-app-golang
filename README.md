# todo-list-app-golang

This is a simple Todo List application built with Golang and Gin framework. It uses PostgreSQL as the database and GORM as the ORM.

## Project Structure
```sh
todo-list-app-golang
├── frontend
│
├── todolist-api/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── models/
│   │   │   └── todo.go
│   │   ├── repository/
│   │   │   └── todo_repository.go
│   │   ├── service/
│   │   │   └── todo_service.go
│   │   └── handlers/
│   │       └── todo_handler.go
│   ├── test/
│   │   └── todo_handler_test.go
│   ├── .env
│   ├── docker-compose.yml
│   ├── go.mod
│   ├── go.sum
│   └── swagger/
```

## Getting Started
### Prerequisites

- Go 1.22.5 or later
- Docker
- PostgreSQL

### Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/todo-list-app-golang.git
cd todo-list-app-golang/todolist-api
```

### Create a .env file in the todolist-api directory with the following content:
```sh
POSTGRES_USER=postgres
POSTGRES_PASSWORD=mysecretpassword
POSTGRES_DB=todolist_db
DB_HOST=localhost
DB_PORT=5432
APP_PORT=8080
```

### Start the PostgreSQL database using Docker:
```sh
docker-compose up -d
```

### Install Go dependencies:
```sh
go mod tidy
```

### Run the application:
```sh
go run cmd/main.go
```

### Api documentation
```sh
http://localhost:8080/swagger/index.html
```

### The application will be running at http://localhost:8080
```sh
API Endpoints
GET /api/v1/todos - Get all todos
POST /api/v1/todos - Create a new todo
```

### Database Migration
```sh
The database migrations are managed using golang-migrate. The migration files are located in the migrations directory.

To run the migrations, the InitDB function in config.go is used, which is called in the main function in main.go.
```

### Project Structure
```sh
cmd/: Contains the main application entry point.
internal/config/: Contains the database configuration and initialization.
internal/handlers/: Contains the HTTP handlers.
internal/models/: Contains the data models.
internal/repository/: Contains the database repository logic.
internal/service/: Contains the business logic.
migrations/: Contains the database migration files.
```

## Deverlop

### Run the swag
When adding a new api you need to update the api documentation (swagger) with the following command
```sh
swag init -g cmd/main.go -o docs
```

### License
This project is licensed under the MIT License - see the LICENSE file for details.

 ```sh
This `README.md` file provides an overview of the project, installation instructions, API endpoints, and a brief description of the project structure.
This `README.md` file provides an overview of the project, installation instructions, API endpoints, and a brief description of the project structure.
```
