# go-connect-albums
Example in Go to create a REST API using Connect

Author Claudio Alarcon clalarco@gmail.com

To start server:

```sh
ALBUMS_DB_TYPE=sqlite DB_PATH=${PWD}/data.sql go run ./src/cmd/server/main.go
```

To start test client:
```sh
go run ./src/cmd/client/main.go
```