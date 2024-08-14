package albums

import (
	"database/sql"
	"errors"

	"clalarco.io/helpers"
)

type db_sqlite struct {
	instance helpers.DbSqlite
}

func (sqlite *db_sqlite) Init() error {
	instance, err := helpers.GetSqlite3Connection()
	if err != nil {
		return err
	}
	sqlite.instance = instance
	schemaErr := sqlite.createSchema()
	if schemaErr != nil {
		return schemaErr
	}
	return nil
}

func (sqlite *db_sqlite) GetAlbums() ([]Album, error) {
	rows, err := sqlite.instance.DB.Query("SELECT * FROM albums")
	if err != nil {
		return []Album{}, err
	}

	var albums []Album = processRows(rows)
	defer rows.Close()
	return albums, nil
}

func (sqlite *db_sqlite) GetAlbum(id string) (Album, error) {
	rows, err := sqlite.instance.DB.Query("SELECT * FROM albums WHERE id = ?", id)
	if err != nil {
		return Album{}, err
	}

	var albums []Album = processRows(rows)
	defer rows.Close()
	if len(albums) == 0 {
		return Album{}, errors.New("ID not found")
	}
	return albums[0], nil
}

func (sqlite *db_sqlite) AddAlbum(album Album) error {
	_, err := sqlite.instance.DB.Exec("INSERT INTO albums VALUES (?, ?, ?, ?)", album.ID, album.Title, album.Artist, album.Price)
	return err
}

func (sqlite *db_sqlite) DeleteAlbum(id string) error {
	_, err := sqlite.instance.DB.Exec("DELETE FROM albums WHERE id = ?", id)
	return err
}

func (sqlite *db_sqlite) createSchema() error {
	_, err := sqlite.instance.DB.Exec("CREATE TABLE IF NOT EXISTS albums (id TEXT PRIMARY KEY, title TEXT, artist TEXT, price REAL)")
	return err
}

func processRows(rows *sql.Rows) []Album {
	var albums []Album
	for rows.Next() {
		var id string
		var title string
		var artist string
		var price float64
		err := rows.Scan(&id, &title, &artist, &price)
		if err != nil {
			return []Album{}
		}
		albums = append(albums, Album{ID: id, Title: title, Artist: artist, Price: price})
	}
	return albums
}
