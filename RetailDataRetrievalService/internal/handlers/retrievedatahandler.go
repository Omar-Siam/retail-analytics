package handlers

import (
	"RetailAnalytics/RetailDataRetrievalService/internal/repository"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
)

func RetrieveDataHandler(s3Client *s3.S3) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		objectKey := r.URL.Query().Get("object")
		data, err := repository.GetObjectFromS3(s3Client, objectKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}
}
