package routes

import (
	"github.com/alias-asso/iosu/plateform/server"
)

func RegisterRoutes(s *server.Server) {
	s.Mux.HandleFunc("/", handleNotFound)
}

// Define all routes
