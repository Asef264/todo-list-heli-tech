package storage

import (
	"context"
	"fmt"
	"log"

	s3Config "github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CreateAWSS3Client(endpoint, accessKey, secretKey string) *s3.Client {
	cfg, err := s3Config.LoadDefaultConfig(context.TODO(), s3Config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}

	cfg.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretKey,
		}, nil
	})
	cfg.BaseEndpoint = aws.String(endpoint)
	fmt.Printf("AWS S3 options -> endpoint: %s, access key: %s, secret key: %s\n", endpoint, accessKey, secretKey)

	client := s3.NewFromConfig(cfg)

	return client
}
