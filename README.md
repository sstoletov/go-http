# Go HTTP Server

A minimal HTTP server written in Go using only the standard library.

## Features

* HTTP server with `net/http`
* REST-style endpoints
* JSON request/response handling
* In-memory user storage
* No third-party frameworks

## Endpoints

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| GET    | `/`           | Home page         |
| GET    | `/users`      | List all users    |
| GET    | `/users?id=1` | Get user by ID    |
| POST   | `/users`      | Create a new user |
| DELETE | `/users?id=1` | Delete a user     |

## Run

```bash
go run main.go
```

The server starts on:

```
http://localhost:8080
```

## Example

```bash
curl http://localhost:8080/users
```

This project was created for learning Go, HTTP, and REST API development without using external frameworks.

