package repository

import (
	"RetailAnalytics/storageservice/internal/models"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

const queueURL = "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/MyLocalQueue"

func SendMessageToSQS(sqsClient *sqs.SQS, payload models.Transaction) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	queueUrl := queueURL
	m, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(payloadBytes)),
		QueueUrl:    &queueUrl,
	})
	if err != nil {
		return err
	}

	log.Printf("SQS MessageID: %v", *m.MessageId)
	return nil
}
