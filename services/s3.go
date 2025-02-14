package services

import (
	"context"
	"time"

	"github.com/IAmRiteshKoushik/mercury/cmd"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client
var PresignClient *s3.PresignClient

func NewS3Clients(cfg aws.Config) {
	S3Client = s3.NewFromConfig(AwsConfig)
	PresignClient = s3.NewPresignClient(S3Client)
}

// For uploading object using presignedURL
func PutPresignURL(key string) string {
	presignedUrl, err := PresignClient.PresignPutObject(context.Background(),
		&s3.PutObjectInput{
			Bucket: aws.String(cmd.EnvVars.AwsBucketName),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(time.Minute*10),
	)
	if err != nil {
		// TODO: Log properly
		return ""
	}
	return presignedUrl.URL
}

func DeleteS3Object(key string) bool {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(cmd.EnvVars.AwsBucketName),
	}

	_, err := S3Client.DeleteObject(context.Background(), input)
	if err != nil {
		return false
	}
	return true
}
