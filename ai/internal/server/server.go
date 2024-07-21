package server

import (
	"log"
	"net/http"

	"github.com/pkg/errors"
)


type ChatClient interface {
	Chat(prompt string) (string, error)
}

type Dependencies struct {
	OpenAIClient ChatClient
}

type Server struct {
	deps *Dependencies
	port string
}

func NewServer(serverDeps *Dependencies) (*Server, error) {
	err := validateServerDependencies(serverDeps)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate server dependencies")
	}


	return &Server{deps: serverDeps}, nil
}

func (s *Server) Start() error {
	router := s.setupRoutes()

	log.Printf("Starting LLM server on port %s...", s.port)
	err := http.ListenAndServe(":8001", router)
	if err != nil {
		return errors.Wrapf(err, "failed to list on port %s", s.port)
	}

	return nil
}

func validateServerDependencies(deps *Dependencies) error {
	if deps.OpenAIClient == nil {
		return errors.New("Server: OpenAI client is required")
	}

	return nil
}
