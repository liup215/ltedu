package ai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"go.uber.org/zap"
)

// Message represents a single chat message used for multi-turn conversations.
type Message struct {
	Role    string // "user", "assistant", or "system"
	Content string
}

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
	// CreateCompletionWithHistory sends a full conversation history to the AI
	// and returns the assistant's next response.
	CreateCompletionWithHistory(messages []Message) (string, error)
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

// CreateCompletionWithHistory sends a full conversation history (including system messages)
// to the AI model and returns the assistant's reply.
func (b *BaseAIClient) CreateCompletionWithHistory(messages []Message) (string, error) {
	params := make([]openai.ChatCompletionMessageParamUnion, 0, len(messages))
	for _, m := range messages {
		switch m.Role {
		case "system":
			params = append(params, openai.SystemMessage(m.Content))
		case "assistant":
			params = append(params, openai.AssistantMessage(m.Content))
		default: // "user"
			params = append(params, openai.UserMessage(m.Content))
		}
	}

	chatCompletion, err := b.client.Chat.Completions.New(
		context.TODO(), openai.ChatCompletionNewParams{
			Messages: params,
			Model:    b.model,
		},
	)

	if err != nil {
		b.logger.Error("Failed to create completion with history", zap.Error(err))
		return "", err
	}

	if len(chatCompletion.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	return chatCompletion.Choices[0].Message.Content, nil
}
