package connector

import (
	"context"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/belary/bedrock/internal/config"
)

type BedrockConnector struct {
	client *bedrockruntime.Client
}

func NewBedrockConnector(cfg *config.Config) (*BedrockConnector, error) {
	awsCfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion(cfg.AWSRegion),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AWSAccessKeyID,
				cfg.AWSSecretAccessKey,
				"",
			),
		),
	)
	if err != nil {
		return nil, err
	}

	client := bedrockruntime.NewFromConfig(awsCfg)
	return &BedrockConnector{client: client}, nil
}

func (bc *BedrockConnector) InvokeModel(ctx context.Context, modelID string, body []byte) ([]byte, error) {
	input := &bedrockruntime.InvokeModelInput{
		ModelId: &modelID,
		Body:    body,
	}

	output, err := bc.client.InvokeModel(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.Body, nil
}
