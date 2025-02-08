package main

import (
	"context"

	"bedrock/internal/config"
	"bedrock/internal/connector"
	"bedrock/internal/models"
	"bedrock/internal/service"
	"bedrock/pkg/utils"
)

func main() {
	// 加载配置

	cfg, err := config.LoadConfig()
	if err != nil {
		utils.ErrorLogger.Fatal(err)
	}

	// 创建 Bedrock 连接器
	bedrockConnector, err := connector.NewBedrockConnector(cfg)
	if err != nil {
		utils.ErrorLogger.Fatal(err)
	}

	// 创建 AI 服务
	aiService := service.NewAIService(bedrockConnector)

	// 示例请求
	request := &models.AIRequest{
		Prompt: "Hello, how are you?",
	}

	// 处理请求
	response, err := aiService.ProcessQuery(context.Background(), request)
	if err != nil {
		utils.ErrorLogger.Printf("Error processing query: %v", err)
		return
	}

	utils.InfoLogger.Printf("Response: %+v", response)
}
