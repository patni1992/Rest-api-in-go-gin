# Complete REST API in Go

https://codingwithpatrik.dev/posts/rest-api-in-gin/

## Development Setup

⚠️ **Important**: Before running the application, make sure to set up the database by following the [Database Migrations](#database-migrations) section below.

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

### Running Without Air

If you prefer not to use Air, you can run the application directly with Go:

```bash
go run ./cmd/api
```

This will start the server on `http://localhost:8080`. Note that you'll need to manually restart the server when you make changes to the code.

### Dependencies

This project uses Go modules for dependency management. Dependencies will be automatically downloaded when you build or run the application. No manual installation is required.

If you want to explicitly download dependencies, you can run:

```bash
go mod download
```

### Database Migrations

This project uses golang-migrate for database migrations. First, install the migrate CLI:

Golang migrate https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md

#### Running Migrations

⚠️ **Required**: Before running the application for the first time, you must run the database migrations to create all necessary database tables and schemas:

```bash
go run ./cmd/migrate up
```

To rollback migrations (undo last migration):

```bash
go run ./cmd/migrate down
```

#### Creating New Migrations

To create a new migration:

```bash
migrate create -ext sql -dir ./cmd/migrate/migrations -seq name_of_migration
```

This will create two new files in the migrations directory:

- `{timestamp}_name_of_migration.up.sql`
- `{timestamp}_name_of_migration.down.sql`

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

To generate the Swagger documentation, run:

```bash
swag init --dir cmd/api --parseDependency --parseInternal --parseDepth 1
```

The API documentation is available via Swagger UI at:

```
http://localhost:8080/swagger
```

This interactive documentation provides:

- Complete API endpoint listing
- Request/response schemas
- Try-it-out functionality
- Authentication details
