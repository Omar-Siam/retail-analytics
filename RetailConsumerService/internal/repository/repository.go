package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
)

func PutObjectToS3(s3Client *s3.S3, bucketName, objectKey string, jsonObject map[string]any) error {
	jsonData, err := json.Marshal(jsonObject)
	if err != nil {
		return err
	}

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		Body:        bytes.NewReader(jsonData),
		ContentType: aws.String(http.DetectContentType(jsonData)),
	})

	if err != nil {
		return err
	}
	fmt.Printf("Object '%s' successfully uploaded to bucket '%s'.\n", objectKey, bucketName)

	return nil
}
