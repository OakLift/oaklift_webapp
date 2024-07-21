package server

import "github.com/gorilla/mux"

func (s *Server) setupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/chat", s.handleCodeReview).Methods("POST")
	return router
}