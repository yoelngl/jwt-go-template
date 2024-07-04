```markdown
# Golang Fiber JWT Authentication Starter

This repository provides a starter template for building a JWT (JSON Web Token) based authentication system using Golang with the Fiber framework and MySQL for the database.

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

### Run database migrations

Ensure your MySQL database is running and configured correctly in the `.env` file. Then run:

```bash
make migrate-up
```

### Start the application

```bash
go run main.go
```

The server will start on `http://localhost:3000`.

## API Endpoints

### Authentication

- **POST /api/v1/auth/register**
  - Register a new user
  - Request body:
    ```json
    {
      "username": "your-username",
      "email": "your-email",
      "password": "your-password"
    }
    ```

- **POST /api/v1/auth/login**
  - Login a user and get a JWT token
  - Request body:
    ```json
    {
      "email": "your-email",
      "password": "your-password"
    }
    ```

### User

- **GET /api/v1/user**
  - Get user information (protected route, requires JWT)
  - Headers:
    ```json
    {
      "Authorization": "Bearer <your-jwt-token>"
    }
    ```

## Running Tests

To run tests, use the following command:

```bash
go test ./...
```

## Makefile Commands

- `make migrate-up` - Run database migrations up
- `make migrate-down` - Rollback database migrations

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://gorm.io/)
- [JWT-Go](https://github.com/dgrijalva/jwt-go)
```
