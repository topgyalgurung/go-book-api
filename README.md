# RESTful Web Service API with Go and Gin

A simple RESTful API built with Go using the Gin Web Framework. This project demonstrates how to route HTTP requests, handle request parameters, and return JSON responses.

**Tutorial Reference:** https://go.dev/doc/tutorial/web-service-gin

## Features

- Build RESTful API endpoints
- Route HTTP requests with Gin
- Parse request parameters
- Return JSON responses
- Manage dependencies using Go modules

## Getting Started

### 1. Initialize a Go module

```bash
go mod init example/go-api
```

Creates a `go.mod` file, which defines the project's module and tracks all dependencies.

### 2. Install the Gin Web Framework

```bash
go get github.com/gin-gonic/gin
```

Downloads and adds the Gin package as a dependency for building the REST API.

### 3. Run the application

```bash
go run main.go
```

Compiles and starts the application. By default, the server runs on:

```
http://localhost:8080
```

## Project Structure

- `main.go` – Entry point of the application
- `go.mod` – Defines the Go module and manages dependencies
- `go.sum` – Stores checksums to ensure dependency integrity


## Testing the API

### Create a Book (POST)

Send a POST request with a JSON request body stored in `body.json`.

```bash

curl http://localhost:8080/books \

  --include \

  --header "Content-Type: application/json" \

  --data @body.json \

  --request POST

```

Example `body.json`:

```json

{

  "id": "4",

  "title": "The Hobbit",

  "author": "J.R.R. Tolkien",

  "price": 12.99

}

```

**Explanation:**

- `--include` (`-i`) displays the response headers.

- `--header "Content-Type: application/json"` tells the server the request body is JSON.

- `--data @body.json` sends the contents of `body.json` as the request body.

- `--request POST` specifies the HTTP POST method.