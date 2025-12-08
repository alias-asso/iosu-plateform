package server

import (
	"net/http"

	"github.com/alias-asso/iosu/plateform/config"
	"github.com/alias-asso/iosu/plateform/database"
	"gorm.io/gorm"
)

// Define a basic http server and connect to the database

type Server struct {
	db          *gorm.DB
	httpServMux *http.ServeMux
}

func NewServer(config config.Config) Server {
	mux := http.NewServeMux()

	err, db := database.ConnectSqlite(config.Sqlite.DbPath)
	if err != nil {

	}
	return Server{
		httpServMux: mux,
		db:          db,
	}
}

func (s *Server) Start(port string) {
	http.ListenAndServe(":"+port, s.httpServMux)
}
