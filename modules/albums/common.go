package albums

import (
	"clalarco.io/helpers"
)

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// ModuleParams holds the parameters for the module.
type ModuleParams struct {
	DBConn string
}

// getModuleParams returns the module parameters.
func getModuleParams() ModuleParams {
	return ModuleParams{
		DBConn: helpers.GetEnv("ALBUMS_DB_TYPE", "mock"),
	}
}

// Database is the interface for a database connection
type DbHandler interface {
	Init() error
	GetAlbums() ([]Album, error)
	GetAlbum(id string) (Album, error)
	AddAlbum(album Album) error
	DeleteAlbum(id string) error
}

// DbFactory returns an instance of IDbConnector based on the provided connectionType.
//
// connectionType: a string representing the type of connection to be used.
// Returns an instance of IDbConnector or nil if the connectionType is not recognized.
func DbFactory(connectionType string) DbHandler {
	var db DbHandler
	switch connectionType {
	case "mock":
		db = &db_mock{}
	case "sqlite":
		db = &db_sqlite{}
	default:
		return nil
	}
	var err = db.Init()
	if err != nil {
		panic(err)
	}
	return db
}

// Handler struct holds required services for handler to function
type Handler struct {
	DB DbHandler
}

func GetHandler() *Handler {
	return &Handler{
		DB: DbFactory(getModuleParams().DBConn),
	}
}
