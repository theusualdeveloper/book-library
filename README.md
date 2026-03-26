# Book Library API

A RESTful API for managing a book library, built with Go using Layered Architecture and Ports & Adapters pattern.

## Architecture

```
├── cmd/api/          # Entry point
├── internal/
│   ├── domain/       # Core business entities (pure Go structs)
│   ├── application/  # Business logic (use cases)
│   ├── port/         # Interfaces (Ports)
│   ├── adapter/      # Interface implementations (Adapters)
│   └── handler/      # HTTP handlers
```

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /books | List all books |
| GET | /books/{id} | Get a book by ID |
| POST | /books | Add a new book |
| PUT | /books/{id} | Update a book |
| DELETE | /books/{id} | Delete a book |

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.21 or higher

### Run

```bash
git clone https://github.com/theusualdeveloper/book-library.git
cd book-library
go run ./cmd/api
```

### Example Requests

```bash
# Add a book
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title":"The Go Programming Language","author":"Alan Donovan","published_year":2015,"genre":"Technology"}'

# List all books
curl http://localhost:8080/books

# Get a book
curl http://localhost:8080/books/{id}

# Update a book
curl -X PUT http://localhost:8080/books/{id} \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Title","author":"Author","published_year":2020,"genre":"Technology"}'

# Delete a book
curl -X DELETE http://localhost:8080/books/{id}
```

## Tech Stack

- **Language:** Go
- **HTTP:** net/http (standard library)
- **Storage:** In-memory

## License

MIT
