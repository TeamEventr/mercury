package files

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/IAmRiteshKoushik/mercury/cmd"
	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

var S3Client *s3.Client
var PresignClient *s3.PresignClient

func GeneratePresignedRequest(ctx context.Context, fileName string,
	validSecs int64) (*v4.PresignedHTTPRequest, error) {
	request, err := PresignClient.PresignPutObject(ctx,
		&s3.PutObjectInput{
			Bucket: aws.String(cmd.EnvVars.CfBucketName),
			Key:    aws.String(fileName),
		}, func(opts *s3.PresignOptions) {
			opts.Expires = time.Duration(validSecs * int64(time.Second))
		})
	if err != nil {
		log.Printf("Could not get a presigned request to put %v:%v. Here's why: %v\n",
			cmd.EnvVars.CfBucketName, fileName, err)
	}
	return request, err
}

func DeleteFile(ctx context.Context, key string) (bool, error) {
	bucket := cmd.EnvVars.CfBucketName
	bypassGovernance := true
	deleted := false
	input := &s3.DeleteObjectInput{
		Bucket:                    aws.String(bucket),
		Key:                       aws.String(key),
		VersionId:                 aws.String(""),
		BypassGovernanceRetention: aws.Bool(bypassGovernance),
	}

	_, err := S3Client.DeleteObject(ctx, input)
	if err != nil {
		var noKey *types.NoSuchKey
		var apiErr *smithy.GenericAPIError
		if errors.As(err, &noKey) {
			log.Printf("Object %s does not exist in %s.\n", key, bucket)
			err = noKey
		} else if errors.As(err, &apiErr) {
			switch apiErr.ErrorCode() {
			case "AccessDenied":
				log.Printf("Access denied: cannot delete object %s from %s.\n", key, bucket)
				err = nil
			case "InvalidArgument":
				if bypassGovernance {
					log.Printf("You cannot specify bypass governance without a lock enabled.")
					err = nil
				}
			}
		}
		return deleted, err
	}

	err = s3.NewObjectNotExistsWaiter(S3Client).Wait(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, time.Minute)
	if err != nil {
		log.Printf("Failed attempt to wait for object %s in bucket %s to be deleted.\n",
			key, bucket)
	}
	deleted = true
	return deleted, err
}
