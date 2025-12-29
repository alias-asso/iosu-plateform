package server

import (
	"html/template"
	"net/http"
)

func (s *Server) handleNotFound(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/pages/error.gohtml"))
	err := tmpl.Execute(w, "404")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
