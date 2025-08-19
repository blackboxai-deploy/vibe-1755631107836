# Go Hello World Server

A simple Go HTTP server that responds with "Hello World" to any request.

## Features

- Responds to any HTTP method (GET, POST, PUT, DELETE, etc.)
- Serves on all routes/paths
- Returns plain text "Hello World" response
- Runs on port 3000

## Running the Server

```bash
# Build and run the server
go run main.go
```

The server will start on `http://localhost:3000`

## API Testing

You can test the server with any HTTP request:

```bash
# GET request
curl http://localhost:3000/

# POST request  
curl -X POST http://localhost:3000/api/test

# Any path works
curl http://localhost:3000/any/path/here
```

All requests will return: `Hello World`

## Building Binary

To build a standalone binary:

```bash
go build -o hello-server main.go
./hello-server
```