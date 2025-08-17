# Recipes

A simple app that can be used to learn the basics of Go, Gin, and GORM.

Swagger: http://localhost:8080/swagger/index.html

## Technologies

- Go
- Gin
- GORM
- Docker Compose

## Run Locally

1. Install Go dependencies:
     ```bash
     go mod tidy  
     ```

2. Run the Go application:
     ```bash
     go run cmd/main.go 
     ```

## Tests & Linting

* Linting:
    ```
    golangci-lint run --fix
    ```

* Unit tests:
    ```
    go test -cover ./...
    ```

## Authors & Copyrights

Shaimaa Sabry

[![GitHub](https://img.icons8.com/ios-glyphs/30/000000/github.png)](https://github.com/ShaimaaSabry)
[![LinkedIn](https://img.icons8.com/ios-filled/30/0A66C2/linkedin.png)](https://www.linkedin.com/in/shaimaa-sabry-161b71a0/)