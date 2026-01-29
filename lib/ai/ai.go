package ai

import (
	"fmt"
	"time"

	client "github.com/aliyun/alibabacloud-bailian-go-sdk/client"
	"go.uber.org/zap"
)

func NewModel(c *Config, logger *zap.Logger) Model {
	if c.Dialect == ALI_BAILIAN {
		tokenClient := &client.AccessTokenClient{
			AgentKey:        &(c.AliBaiLian.AgentKey),
			AccessKeyId:     &(c.AliBaiLian.AccessKey),
			AccessKeySecret: &(c.AliBaiLian.AccessSecretKey),
		}

		return &Bailian{
			tokenClient: tokenClient,
			logger:      logger,
			appId:       c.AliBaiLian.AppId,
		}
	}
	return nil
}

type Model interface {
	CreateCompletion(string) (string, error)
}

type Bailian struct {
	tokenClient *client.AccessTokenClient
	appId       string
	logger      *zap.Logger
}

func (b *Bailian) CreateCompletion(prompt string) (string, error) {
	token, err := b.tokenClient.GetToken()
	if err != nil {
		return "", err
	}

	cc := client.CompletionClient{Token: &token}
	cc.SetTimeout(15 * time.Second)

	request := &client.CompletionRequest{}
	request.SetAppId(b.appId)
	request.SetPrompt(prompt)

	response, err := cc.CreateCompletion(request)
	if err != nil {
		b.logger.Error("%v\n", zap.Error(err))
		return "", err
	}

	if !response.Success {
		estr := fmt.Sprintf("failed to create completion, requestId: %s, code: %s, message: %s\n", response.GetRequestId(), response.GetCode(), response.GetMessage())
		b.logger.Info(estr)
		return "", err
	}

	return response.GetData().GetText(), nil
}
