package albums

import "errors"

// Fixtures for album REST api
type db_mock struct{}

var mockAlbums = map[string]Album{
	"1": {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func (m *db_mock) Init() error {
	return nil
}

func (m *db_mock) GetAlbums() ([]Album, error) {
	var albums []Album
	for _, a := range mockAlbums {
		albums = append(albums, a)
	}
	return albums, nil
}

func (m *db_mock) GetAlbum(id string) (Album, error) {
	a, found := mockAlbums[id]
	if !found {
		return Album{}, errors.New("ID not found")
	}
	return a, nil
}

func (m *db_mock) AddAlbum(album Album) error {
	_, found := mockAlbums[album.ID]
	if !found {
		return errors.New("ID already exists")
	}
	mockAlbums[album.ID] = album
	return nil
}

func (m *db_mock) DeleteAlbum(id string) error {
	_, found := mockAlbums[id]
	if !found {
		return errors.New("ID not found")
	}
	delete(mockAlbums, id)
	return nil
}
