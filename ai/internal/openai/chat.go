package openai

import (
	"context"
	"time"

	"github.com/pkg/errors"
	openailib "github.com/sashabaranov/go-openai"
)

func (c *Client) CodeReview(prompt string) (string, error) {
	return c.sendPrompt(prompt)
}

func (c *Client) CreateProject(prompt string) (string, error) {
	return c.sendPrompt(prompt)
}

func (c *Client) sendPrompt(prompt string) (string, error)  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ccm := openailib.ChatCompletionMessage{
		Role: openailib.ChatMessageRoleUser,
		Content: prompt,
	}

	req := openailib.ChatCompletionRequest{
		Model: openailib.GPT3Dot5Turbo,
		Messages: []openailib.ChatCompletionMessage{ccm},
	}

	resp, err := c.openAIClient.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", errors.Wrap(err, "OpenAI/Client: unable to send chat completion request to OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}
