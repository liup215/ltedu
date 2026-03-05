package ai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"go.uber.org/zap"
)

// Message represents a single chat message with a role and content.
type Message struct {
	Role    string // "system", "user", or "assistant"
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

// Model is the interface for AI language model clients.
type Model interface {
	// CreateCompletion sends a single user prompt and returns the model response.
	CreateCompletion(string) (string, error)
	// CreateCompletionWithMessages sends a multi-turn message list (supporting system
	// prompts and conversation history) and returns the model response.
	CreateCompletionWithMessages(messages []Message) (string, error)
}

type BaseAIClient struct {
	client *openai.Client
	model  string
	logger *zap.Logger
}

func (b *BaseAIClient) CreateCompletion(prompt string) (string, error) {
	return b.CreateCompletionWithMessages([]Message{{Role: "user", Content: prompt}})
}

func (b *BaseAIClient) CreateCompletionWithMessages(messages []Message) (string, error) {
	params := make([]openai.ChatCompletionMessageParamUnion, 0, len(messages))
	for _, m := range messages {
		switch m.Role {
		case "system":
			params = append(params, openai.SystemMessage(m.Content))
		case "assistant":
			params = append(params, openai.AssistantMessage(m.Content))
		default:
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
		b.logger.Error("Failed to create completion", zap.Error(err))
		return "", err
	}

	if len(chatCompletion.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	return chatCompletion.Choices[0].Message.Content, nil
}
