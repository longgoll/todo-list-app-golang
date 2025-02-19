# todo-list-app-golang

This is a simple Todo List application built with Golang and Gin framework. It uses PostgreSQL as the database and GORM as the ORM.

## Project Structure
```sh
todolist-api/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── models/
│   │   └── todo.go
│   ├── repository/
│   │   └── todo_repository.go
│   ├── service/
│   │   └── todo_service.go
│   └── handlers/
│       └── todo_handler.go
├── test/
│   └── todo_handler_test.go
├── .env
├── docker-compose.yml
├── go.mod
├── go.sum
└── swagger/
```

## Getting Started

frontend/ todolist-api/ .env cmd/ main.go docker-compose.yml go.mod go.sum internal/ config/ config.go handlers/ todo_handler.go models/ todo.go repository/ todo_repository.go service/ todo_service.go migrations/ 0001_create_todos_table.down.sql 0001_create_todos_table.up.sql swagger/ test/

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

File `README.md` này cung cấp một cái nhìn tổng quan về dự án, hướng dẫn cài đặt, các endpoint của API, và mô tả ngắn gọn về cấu trúc dự án.
File `README.md` này cung cấp một cái nhìn tổng quan về dự án, hướng dẫn cài đặt, các endpoint của API, và mô tả ngắn gọn về cấu trúc dự án.# todo-list-app-golang
```
