![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![SQL](https://img.shields.io/badge/SQL-database-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white)

## Description

Written in GO, cuz trying to learn the language.

A simple backend that executes raw SQL Queries.

A UI to actually input the queries and see the results

DB: postgres

### Setup

before starting, run:
```
docker run --name go-db -e POSTGRES_PASSWORD=complex_and_hard_password -p 5432:5432 -d postgres
```

installing a postgres driver:
```
go mod init test-sql
go get github.com/lib/pq
```
(go.mod & go.sum will be created)

then run:
`go run main.go`

and go to `http://localhost:8080`
