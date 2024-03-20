package openaix

import (
	"github.com/blackhorseya/mundo/pkg/configx"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"github.com/sashabaranov/go-openai"
)

const (
	defaultMaxTokens   = 300
	defaultModel       = openai.GPT3Dot5Turbo
	defaultTemperature = 1.0
	defaultTopP        = 1.0
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

// CreateFunctionCall is used to create a function call.
func (c *Client) CreateFunctionCall(
	ctx contextx.Contextx,
	content string,
) (resp openai.ChatCompletionResponse, err error) {
	return c.Client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: defaultModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
		},
		MaxTokens:        defaultMaxTokens,
		Temperature:      defaultTemperature,
		TopP:             defaultTopP,
		PresencePenalty:  0.0,
		FrequencyPenalty: 0.0,
	})
}
