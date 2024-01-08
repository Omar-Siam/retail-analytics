package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

const sqsQueueURL = "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/MyLocalQueue"

func PollSQS(sqsClient *sqs.SQS) ([]*sqs.Message, error) {
	result, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(sqsQueueURL),
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(20),
	})
	if err != nil {
		return nil, err
	}
	return result.Messages, nil
}

func DeleteMessageFromSQS(sqsClient *sqs.SQS, msg *sqs.Message) {
	_, err := sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(sqsQueueURL),
		ReceiptHandle: msg.ReceiptHandle,
	})
	if err != nil {
		log.Printf("Failed to delete message from SQS: %v", err)
	}
}
