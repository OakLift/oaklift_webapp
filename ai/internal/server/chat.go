package server

import (
	"io"
	"net/http"

	"ai-tooling.com/internal/reviewer"
)

func (s *Server) handleCodeReview(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	promptString := string(bodyBytes)
	finalPrompt := reviewer.SanitizePrompt(promptString)

	response, err := s.deps.OpenAIClient.Chat(finalPrompt)
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