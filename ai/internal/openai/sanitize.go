package openai

import (
	"fmt"
)

func SanitizePrompt(prompt string) string {
	sanitizedPrompt := formAsQuestion(prompt)
	return sanitizedPrompt
}

func formAsQuestion(prompt string) string {
	const question = "Can you help review the following code snippet and tell me how I can improve it?"
	newPrompt := fmt.Sprintf("%s\n\n%s", question, prompt)
	return newPrompt
}
