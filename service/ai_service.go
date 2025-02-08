package service

import (
	"context"
	"encoding/json"

	"github.com/belary/bedrock/internal/connector"
	"github.com/belary/bedrock/internal/models"
)

type AIService struct {
	connector *connector.BedrockConnector
}

func NewAIService(connector *connector.BedrockConnector) *AIService {
	return &AIService{
		connector: connector,
	}
}

func (s *AIService) ProcessQuery(ctx context.Context, request *models.AIRequest) (*models.AIResponse, error) {
	// 准备请求体
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// 调用模型
	response, err := s.connector.InvokeModel(ctx, "anthropic.claude-v2", requestBody)
	if err != nil {
		return &models.AIResponse{Error: err.Error()}, err
	}

	// 解析响应
	var aiResponse models.AIResponse
	if err := json.Unmarshal(response, &aiResponse); err != nil {
		return nil, err
	}

	return &aiResponse, nil
}
