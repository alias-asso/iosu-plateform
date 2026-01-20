package server

func registerRoutes(s *Server) {
	s.mux.HandleFunc("/", s.handleNotFound)

	// Login routes
	s.mux.HandleFunc("GET /login", s.getLogin)
	s.mux.HandleFunc("POST /login", s.postLogin)
	s.mux.HandleFunc("POST /register", withAdmin(s.postRegisterAccount))
}
