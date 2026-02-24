package ai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"go.uber.org/zap"
)

func NewModel(c *Config, logger *zap.Logger) Model {
	client := openai.NewClient(
		option.WithAPIKey(c.ApiKey),
		option.WithBaseURL(c.BaseUrl),
	)

	return &BaseAIClient{
		client: &client,
		model:  c.Model,
		logger: logger,
	}
}

type Model interface {
	CreateCompletion(string) (string, error)
}

type BaseAIClient struct {
	client *openai.Client
	model  string
	logger *zap.Logger
}

func (b *BaseAIClient) CreateCompletion(prompt string) (string, error) {
	chatCompletion, err := b.client.Chat.Completions.New(
		context.TODO(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(prompt),
			},
			Model: b.model,
		},
	)

	if err != nil {
		b.logger.Error("Failed to create completion", zap.Error(err))
		return "", err
	}

	if len(chatCompletion.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	return chatCompletion.Choices[0].Message.Content, nil
}
