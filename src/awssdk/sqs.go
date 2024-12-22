package awssdk

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSClient struct {
	client   *sqs.Client
	queueURL *string
}

func NewSQSClient(ctx context.Context, cfg aws.Config, queueName string) (*SQSClient, error) {
	client := sqs.NewFromConfig(cfg)

	// Get queue URL from queue name
	result, err := client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		return nil, err
	}

	return &SQSClient{
		client:   client,
		queueURL: result.QueueUrl,
	}, nil
}

func (c *SQSClient) SendMessage(ctx context.Context, messageBody string) error {
	input := &sqs.SendMessageInput{
		QueueUrl:    c.queueURL,
		MessageBody: aws.String(messageBody),
	}

	_, err := c.client.SendMessage(ctx, input)
	if err != nil {
		log.Printf("Error sending message to SQS: %v", err)
		return err
	}

	return nil
}
