package openaix

import (
	"github.com/blackhorseya/mundo/pkg/configx"
	"github.com/sashabaranov/go-openai"
)

// Client is the client for openai.
type Client struct {
	*openai.Client
}

// NewClient is used to create a new openai client.
func NewClient() (*Client, error) {
	client := openai.NewClient(configx.C.OpenAI.Key)

	return &Client{
		Client: client,
	}, nil
}
