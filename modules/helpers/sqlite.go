package helpers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DbSqlite struct {
	DB *sql.DB
}

var _db_instance = DbSqlite{}

// GetConnection returns a connection to the SQLite database.
//
// It checks if the database connection is already established and returns it if it is.
// Otherwise, it opens a new connection to the database located at "./data/albums.db" and returns it.
//
// Returns:
// - DbSqlite: the database connection.
// - error: an error if there was a problem opening the database connection.
func GetSqlite3Connection() (DbSqlite, error) {
	if _db_instance.DB == nil {
		database, err := sql.Open("sqlite3", GetEnv("DB_PATH", "sqlite.db"))
		if err != nil {
			return DbSqlite{}, err
		}
		_db_instance.DB = database
	}
	return _db_instance, nil
}
