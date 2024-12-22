package factory

import (
	"context"
	"fmt"

	"github.com/yoelsusanto/the-workshop/src/awssdk"
)

const (
	QueueMain = "main"
)

func newSQSClient(ctx context.Context, queueName string) *awssdk.SQSClient {
	cfg := CreateAWSConfig(ctx)

	sqsClient, err := awssdk.NewSQSClient(ctx, cfg, queueName)
	if err != nil {
		panic(fmt.Errorf("factory: failed to create SQS client: %w", err))
	}

	return sqsClient
}

func NewMainQueue(ctx context.Context) *awssdk.SQSClient {
	return newSQSClient(ctx, QueueMain)
}
