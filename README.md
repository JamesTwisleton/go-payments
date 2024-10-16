# go-payments

This is a simple Go application that provides a RESTful API to generate random credit scores. The project uses Go modules for dependency management. It was created following [this](https://medium.com/@Moesif/building-a-restful-api-with-go-dbd6e7aecf87) tutorial.

## Prerequisites

### Running the app
- [Go](https://golang.org/dl/)

### Running the pre-commit hooks
- [golangci-lint](https://github.com/golangci/golangci-lint)

## Getting Started

To get a local instance up and running, follow these steps.

### 1. Download dependencies
Navigate to the repository root folder and run:
```bash
go mod tidy
```

### 2. Run the application
```bash
go run .
```

This command will start an HTTP server on port 8080. The server will handle requests to the `/creditscore` endpoint, returning a random credit score. 

### 3. Access the API
Visit [http://localhost:8080/creditscore](http://localhost:8080/creditscore) in your browser. The endpoint will return a random credit score in JSON format. Stop the server with `CTRL + C`

## Pre-commit Hooks

This project includes pre-commit hooks to ensure code quality and consistency before any changes are committed.

### What do the pre-commit hooks do?

- **Code Formatting**: Runs `gofmt` to automatically format the code.
- **Linting**: Runs `golangci-lint` to check for common code issues and enforce Go best practices.
- **Build Check**: Ensures that the project compiles successfully before committing.

### How to Set Up the Pre-commit Hook


Copy the pre-commit script to your local [.git/hooks](.git/hooks) directory and make it executable:


```bash
cp scripts/pre-commit .git/hooks/
chmod +x .git/hooks/pre-commit
```
