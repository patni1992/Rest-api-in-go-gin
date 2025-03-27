# Complete REST API in Go

https://codingwithpatrik.dev/posts/rest-api-in-gin/

## Development Setup

### Using Air for Live Reload

This project uses Air for live reload during development. Here's how to get started:

1. Install Air:

```bash
go install github.com/cosmtrek/air@latest
```

2. Run the application with Air:

```bash
air
```

Air will automatically:

- Watch your Go files for changes
- Rebuild your application when files change
- Restart your server automatically
- Show build errors in a colorized format

To stop Air, press `Ctrl+C` in your terminal.

The project is already configured with a `.air.toml` file that:

- Watches the `cmd/api` directory
- Excludes test files and common directories
- Includes Go files, templates, and HTML files
- Uses colorized output for better visibility

### Building the Application

To build the application:

```bash
go build -o api ./cmd/api
```

This will create an executable named `api` in your project root directory.

### Running the Application

After building, you can run the application using:

```bash
./api
```

The server will start on `http://localhost:8080` by default.

### API Documentation

The API documentation is available via Swagger UI at:

```
http://localhost:8080/swagger
```

This interactive documentation provides:

- Complete API endpoint listing
- Request/response schemas
- Try-it-out functionality
- Authentication details
