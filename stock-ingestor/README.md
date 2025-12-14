# Stock Ingestor

A Go application that ingests stock data from Alpha Vantage API and serves it via a REST API using Fiber framework with CORS support.

## Architecture

This project follows Clean Architecture principles:

- **Domain Layer**: Contains entities and repository interfaces.
- **Application Layer**: Contains use cases and DTOs.
- **Infrastructure Layer**: Contains external API adapters and implementations.
- **Interface Layer**: Contains HTTP handlers and middleware.

## Project Structure

```
.
├── cmd/
│   └── api/                 # Application entry point
├── config/                  # Configuration loading
├── internal/
│   ├── application/
│   │   ├── dto/            # Data Transfer Objects
│   │   └── usecases/       # Application business logic
│   ├── domain/
│   │   ├── entities/       # Domain entities
│   │   └── repositories/   # Repository interfaces
│   ├── infrastructure/
│   │   └── external/
│   │       └── alphavantage/  # External API adapter
│   └── interfaces/
│       └── http/
│           ├── handlers/   # HTTP handlers
│           └── middleware/ # HTTP middleware
├── pkg/                     # Shared packages
└── config.yaml              # Configuration file
```

## API Endpoints

- `GET /api/stocks/:symbol?period=weekly` - Get stock time series data

## Running the Application

1. Ensure you have a valid Alpha Vantage API key.
2. Update `config.yaml` with your API key.
3. Run `go run main.go`
4. The server will start on port 3000.

## Dependencies

- [Fiber](https://github.com/gofiber/fiber) - Web framework
- [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3) - YAML parser