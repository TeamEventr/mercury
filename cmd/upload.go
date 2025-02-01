package cmd

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	bucketName string
	keyName    string
)

// func newAwsConfig() aws.Config {
// 	cfg, err := config.LoadDefaultConfig(
// 		context.Background(),
// 		config.WithRegion(regionName),
// 		config.WithSharedConfigProfile(profileName),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return cfg
// }

// For downloading object using presignedURL
func getPresignedURL(cfg aws.Config) string {
	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)
	presignedUrl, err := presignClient.PresignGetObject(context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(keyName),
		},
		s3.WithPresignExpires(time.Minute*10),
	)
	if err != nil {
		log.Fatal(err)
	}
	return presignedUrl.URL
}

// For uploading object using presignedURL
func putPresignURL(cfg aws.Config) string {
	// FIXME: Bad object path name (dummy text)
	objectName := "myobject"

	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)
	presignedUrl, err := presignClient.PresignPutObject(context.Background(),
		&s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectName),
		},
		s3.WithPresignExpires(time.Minute*15),
	)
	if err != nil {
		log.Fatal(err)
	}
	return presignedUrl.URL
}

// Locally testing file upload
// Would not be used in production
func uploadFile(filePath string, url string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	buffer := bytes.NewBuffer(nil)

	if _, err := io.Copy(buffer, file); err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, url, buffer)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "multipart/form-data")
	client := &http.Client{}
	_, err = client.Do(request)
	return err
}
