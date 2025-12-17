package server

import (
	"log"
	"net/http"

	"github.com/alias-asso/iosu/plateform/config"
	"github.com/alias-asso/iosu/plateform/database"
	"gorm.io/gorm"
)

type Server struct {
	Db  *gorm.DB
	Mux *http.ServeMux
}

// Define a basic http server and connect to the database
func NewServer(config config.Config) (Server, error) {
	mux := http.NewServeMux()

	err, db := database.ConnectSqlite(config.Sqlite.DbPath)
	if err != nil {
		return Server{}, err
	}
	return Server{
		Mux: mux,
		Db:  db,
	}, nil
}

func (s *Server) Start(port string) {
	log.Printf("Listening on %s:%s", "localhost", port)
	http.ListenAndServe(":"+port, s.Mux)
}
