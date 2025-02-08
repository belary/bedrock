package service

import (
	"context"
	"encoding/json"

	"github.com/belary/bedrock/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type BedrockClient struct {
	client *bedrockruntime.Client
	config *config.Config
}

type ModelRequest struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
}

func NewBedrockClient(cfg *config.Config) (*BedrockClient, error) {
	// 加载默认 AWS 配置
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.AWSRegion),
	)
	if err != nil {
		return nil, err
	}

	// 创建 Bedrock Runtime 客户端
	client := bedrockruntime.NewFromConfig(awsCfg)

	return &BedrockClient{
		client: client,
		config: cfg,
	}, nil
}

func (c *BedrockClient) InvokeModel(prompt string) (string, error) {
	request := ModelRequest{
		Prompt:      prompt,
		MaxTokens:   c.config.MaxTokens,
		Temperature: c.config.Temperature,
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	input := &bedrockruntime.InvokeModelInput{
		ModelId:     aws.String(c.config.ModelID),
		ContentType: aws.String("application/json"),
		Accept:      aws.String("application/json"),
		Body:        payload,
	}

	output, err := c.client.InvokeModel(context.TODO(), input)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(output.Body, &response); err != nil {
		return "", err
	}

	// 根据实际响应结构进行解析
	// 这里需要根据具体使用的模型来调整响应解析方式
	return response["completion"].(string), nil
}
