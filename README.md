# Self-Hosted HIBP API with Go

This repository demonstrates a self-hosted [Have I Been Pwned (HIBP)](https://haveibeenpwned.com/Passwords) API using Go and the standard library. The project includes a script to convert a hibp.txt file (containing password hashes and their counts) into a SQLite database and an HTTP API to query the data. this repo source including:

1. **Convert HIBP Text File to SQLite:**
   A script to convert the `hibp.txt` file into a SQLite database for fast and efficient queries.

2. **API for Hash Query:**
   A simple HTTP API to check if a password hash prefix and suffix exist in the database.

3. **Go Standard Library:**
   the project uses only the Go standard library for simplicity and performance.

Project Structure:

├── cmd
│ └── api
│ └── main.go # API server entry point
├── internal
│ ├── database
│ │ └── database.go # Database initialization and queries
│ └── handler
│ └── handler.go # API request handlers
├── scripts
│ └── text_to_sqlite.go # Script to convert hibp.txt to SQLite
├── data
│ └── hibp.db # Generated SQLite database (ignored in .gitignore)
├── .env # Environment variables
├── go.mod # Go module file
└── README.md # Project documentation

## Run the script to convert the hibp.txt file to a SQLite database

```bash
go run scripts/text_to_sqlite.go
```

## Run the server

```bash
go run cmd/api/main.go
```

## env variable

```
SERVER_PORT=3000
DB_PATH=../../data/hbip.db
TEXT_PATH=../../data/hbip_example.txt
API_VERSION=v1
```

## Performance Considerations

- Compression: You can add Gzip compression to reduce the response size and improve performance.
- Caching: Add caching headers to reduce repeated queries for the same hash prefix.

## License

This project is licensed under the [MIT License](https://opensource.org/license/mit). See the LICENSE file for details.
