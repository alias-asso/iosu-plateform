package server

func registerRoutes(s *Server) {
	s.mux.HandleFunc("/", s.getNotFound)

	// Login routes
	s.mux.HandleFunc("GET /login", s.getLogin)
	s.mux.HandleFunc("POST /login", s.postLogin)

	// Register routes
	s.mux.HandleFunc("POST /register", withAdmin(s.postRegisterAccount))
	s.mux.HandleFunc("POST /register/batch", withAdmin(s.postBatchCreateAccounts))
}
