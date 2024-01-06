package clients

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func NewS3Client() *s3.S3 {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("minioaccesskey", "miniosecretkey", ""),
		Endpoint:         aws.String("http://localhost:9000"),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}

	return s3.New(newSession)
}
