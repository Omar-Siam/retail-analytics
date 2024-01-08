package clients

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

func NewSQSClient() *sqs.SQS {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("mock_access_key", "mock_secret_key", ""),
		Endpoint:         aws.String("http://localhost:4566"),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}

	return sqs.New(newSession)
}
