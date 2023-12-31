# Go Banking App

This is a simple banking application written in Go. It uses SQLC for generating type safe Go from SQL.

This toy project is to learn:

1. Concurrency in Go
2. Database Operations
3. Test-driven Development
4. Error Handling
5. Context Usage
6. Docker Usage

## Project Structure

- `db/` contains SQL migrations and queries.
- `sqlc/` contains Go code generated by SQLC from the SQL queries.
- `util/` contains utility functions used across the application.
- `go.mod` and `go.sum` are Go modules files.
- `docker-compose.yml` is used to run the application in a Docker container.

## Key Files

- `db/sqlc/db.go`: This is the main database file where the database connection is established.
- `sqlc/account.sql.go`: Contains Go code for account related SQL queries.
- `sqlc/store.go`: Contains the main store functions for the application.
- `sqlc/store_test.go`: Contains tests for the store functions.
- `util/random.go`: Contains utility functions for generating random data.

## How to Run

You can use Docker Compose to run setup the dev environment. Simply run the following command:

```sh
docker-compose up
```

Testing
Tests are written using the testing package from the Go standard library. You can run the tests with the following command:
```sh
go test ./...
```
