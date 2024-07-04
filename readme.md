---

# Golang Fiber JWT Authentication Starter

This documentation provides a comprehensive guide for setting up and using the JWT authentication starter template built with Golang, Fiber, and MySQL.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Environment Variables](#environment-variables)
- [Database Migrations](#database-migrations)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
  - [Authentication](#authentication)
  - [User](#user)
- [Testing](#testing)
- [Makefile Commands](#makefile-commands)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Introduction

This project is a starter template for building a JWT-based authentication system using Golang with the Fiber framework and MySQL as the database. It provides a basic setup for user registration, login, and protected routes using JWT tokens.

## Features

- User registration and login
- JWT-based authentication and authorization
- Middleware for protected routes
- Database migrations
- Structured project layout

## Project Structure

```
.
├── api
│   ├── handler
│   │   ├── auth_handler.go
│   │   └── user_handler.go
│   ├── middleware
│   │   └── auth_middleware.go
│   └── router
│       └── router.go
├── config
│   └── config.go
├── internal
│   ├── repository
│   │   ├── auth_repository.go
│   │   └── user_repository.go
│   └── service
│       ├── auth_service.go
│       └── user_service.go
├── migrations
│   ├── 20240604124706_create_users_table.down.sql
│   └── 20240604124706_create_users_table.up.sql
├── model
│   ├── jwt.go
│   ├── response.go
│   └── user.go
├── pkg
│   ├── constant
│   │   ├── auth_message.go
│   │   ├── http_message.go
│   │   └── user_message.go
│   └── utils
│       ├── jwt.go
│       └── response.go
├── .env.example
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── Makefile
```

## Prerequisites

- Golang installed (version 1.16+)
- MySQL installed and running
- Go modules enabled

## Getting Started

### Clone the repository

```bash
git clone https://github.com/your-username/golang-fiber-jwt-auth-starter.git
cd golang-fiber-jwt-auth-starter
```

### Setup environment variables

Copy the example environment file and modify the variables as needed:

```bash
cp .env.example .env
```

### Install dependencies

```bash
go mod tidy
```

## Environment Variables

The environment variables required for the application are defined in the `.env` file. Here is an example of the variables you need to set:

```
APP_PORT=3000
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
JWT_SECRET=yourjwtsecret
```

## Database Migrations

To set up the database schema, run the following command:

```bash
make migrate-up
```

To rollback the migrations, run:

```bash
make migrate-down
```

## Running the Application

To start the application, run:

```bash
go run main.go
```

The server will start on `http://localhost:3000`.

## API Endpoints

### Authentication

#### Register

- **Endpoint**: `POST /api/v1/auth/register`
- **Description**: Register a new user
- **Request Body**:
  ```json
  {
    "username": "your-username",
    "email": "your-email",
    "password": "your-password"
  }
  ```
- **Response**:
  - `201 Created`: User successfully registered
  - `400 Bad Request`: Invalid request body

#### Login

- **Endpoint**: `POST /api/v1/auth/login`
- **Description**: Login a user and get a JWT token
- **Request Body**:
  ```json
  {
    "email": "your-email",
    "password": "your-password"
  }
  ```
- **Response**:
  - `200 OK`: Successfully authenticated, returns JWT token
  - `401 Unauthorized`: Invalid credentials

### User

#### Get User Information

- **Endpoint**: `GET /api/v1/user`
- **Description**: Get user information (protected route)
- **Headers**:
  ```json
  {
    "Authorization": "Bearer <your-jwt-token>"
  }
  ```
- **Response**:
  - `200 OK`: Returns user information
  - `401 Unauthorized`: Invalid or missing JWT token

## Testing

To run tests, use the following command:

```bash
go test ./...
```

## Makefile Commands

- `make migrate-up`: Run database migrations up
- `make migrate-down`: Rollback database migrations

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://gorm.io/)
- [JWT-Go](https://github.com/dgrijalva/jwt-go)

---
