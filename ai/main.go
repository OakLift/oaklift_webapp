package main

import (
	"fmt"
	"log"
	"os"

	"ai-tooling.com/internal/openai"
	"ai-tooling.com/internal/server"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	fmt.Println(apiKey)
	openAIClient, err := openai.NewClient()
	if err != nil {
		log.Fatalf("OpenAI: failed to construct new client: %v", err)
	}

	deps := &server.Dependencies{OpenAIClient: openAIClient}
	s, err := server.NewServer(deps)
	if err != nil {
		log.Fatalf("CRITICAL: failed to construct LLM server: %v", err)
	}

	err = s.Start()
	if err != nil {
		log.Fatalf("CRITICAL: failed to start LLM server: %v", err)
	}
}
