package server

import "github.com/gorilla/mux"

const (
	codeReviewPath = "/code/review"
	projectCreatePath = "/project/create"
)

func (s *Server) setupRoutes() *mux.Router {
	router := mux.NewRouter()
	v1Router := router.PathPrefix("/api/v1").Subrouter()
	v1Router.HandleFunc(codeReviewPath, s.handleCodeReview).Methods("POST")
	v1Router.HandleFunc(projectCreatePath, s.handleProjectCreate).Methods("POST")
	
	return router
}
