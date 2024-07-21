package main

import (
	"log"

	"ai-tooling.com/internal/openai"
	"ai-tooling.com/internal/server"
)



func main() {
	openAIClient, err := openai.NewClient()
	if err != nil {
		log.Fatal("OpenAI: failed to construct new client")
	}

	deps := &server.Dependencies{
		OpenAIClient: openAIClient,
	}

	s, err := server.NewServer(deps)
	if err != nil {
		log.Fatalf("CRITICAL: failed to construct LLM server: %v", err)
	}

	err = s.Start()
	if err != nil {
		log.Fatalf("CRITICAL: failed to start LLM server: %v", err)
	}
}
