# todo-list-app-golang

This is a simple Todo List application built with Golang and Gin framework. It uses PostgreSQL as the database and GORM as the ORM.

## Project Structure

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

POSTGRES_USER=postgres
POSTGRES_PASSWORD=mysecretpassword
POSTGRES_DB=todolist_db
DB_HOST=localhost
DB_PORT=5432
APP_PORT=8080

docker-compose up -d

go mod tidy

go run cmd/main.go



File `README.md` này cung cấp một cái nhìn tổng quan về dự án, hướng dẫn cài đặt, các endpoint của API, và mô tả ngắn gọn về cấu trúc dự án.
File `README.md` này cung cấp một cái nhìn tổng quan về dự án, hướng dẫn cài đặt, các endpoint của API, và mô tả ngắn gọn về cấu trúc dự án.# todo-list-app-golang
