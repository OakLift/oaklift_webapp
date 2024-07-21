package openai

import (
	"github.com/pkg/errors"
	openailib "github.com/sashabaranov/go-openai"
)

var (
	ErrNilClient = errors.New("nil client")
)

type Client struct {
	openAIClient *openailib.Client
}

func NewClient() (*Client, error){
	cfg, err := NewConfig()
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize OpenAI client configuration")
	}

	client := openailib.NewClient(cfg.APIKey)
	if client == nil {
		return nil, ErrNilClient
	}
	return &Client{openAIClient: client}, nil
}