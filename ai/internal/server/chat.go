package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"ai-tooling.com/internal/openai"
	"ai-tooling.com/internal/openai/project-creator"
)

func (s *Server) handleCodeReview(w http.ResponseWriter, r *http.Request) {
	log.Println("handleCodeReview handler reached")

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	promptString := string(bodyBytes)
	finalPrompt := openai.SanitizePrompt(promptString)

	response, err := s.deps.OpenAIClient.CodeReview(finalPrompt)
	if err != nil {
		http.Error(w, "failed to chat with OpenAI", http.StatusInternalServerError)
		return
	}

	if response == "" {
		http.Error(w, "OpenAI response was empty", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (s *Server) handleProjectCreate(w http.ResponseWriter, r *http.Request) {
	log.Println("handleProjectCreate handler reached")

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	createRequest := &project.CreateRequest{}
	err = json.Unmarshal(bodyBytes, createRequest)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to unmarshal request body", http.StatusBadRequest)
		return
	}

	finalPrompt, err := project.CreatePrompt(createRequest.Language, createRequest.Level)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create project", http.StatusInternalServerError)
		return
	}

	response, err := s.deps.OpenAIClient.CreateProject(finalPrompt)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to chat with OpenAI", http.StatusInternalServerError)
		return
	}

	if response == "" {
		log.Println("empty response")
		http.Error(w, "OpenAI response was empty", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
