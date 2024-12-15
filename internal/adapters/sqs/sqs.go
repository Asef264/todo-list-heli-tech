package sqs_adapter

import (
	"context"
	"log"
	"todo-list/config"
	ports "todo-list/internal/ports/sqs"
	sqsPkg "todo-list/pkg/sqs"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type sqsAdapter struct {
	sqsAction sqsPkg.SQSClient
}

func NewSQSAdapter(sqsAction sqsPkg.SQSClient) ports.SQS {
	return &sqsAdapter{
		sqsAction: sqsAction,
	}
}

func (sa sqsAdapter) SendMessage(ctx context.Context, message string) error {
	_, err := sa.sqsAction.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: &message,
		QueueUrl:    &config.AppConfig.SQSConfig.QueueURL,
	}, nil)
	if err != nil {
		log.Printf("Failed to send message to SQS: %v", err)
	}
	log.Println("message successfully sent")
	return nil
}
