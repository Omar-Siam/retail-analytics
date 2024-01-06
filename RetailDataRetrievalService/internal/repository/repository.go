package repository

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

const bucketName = "test-s3-bucket"

func GetObjectFromS3(s3Client *s3.S3, objectKey string) (map[string]any, error) {
	resp, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonObject map[string]any
	err = json.Unmarshal(body, &jsonObject)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Object '%s' successfully retrieved from bucket '%s'.\n", objectKey, bucketName)
	return jsonObject, nil
}
