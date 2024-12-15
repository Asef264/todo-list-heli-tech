package storage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// func CreateAWSS3Client(endpoint, accessKey, secretKey string) *s3.Client {
// 	cfg, err := s3Config.LoadDefaultConfig(context.TODO(), s3Config.WithRegion("me-south-1"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	cfg.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
// 		return aws.Credentials{
// 			AccessKeyID:     accessKey,
// 			SecretAccessKey: secretKey,
// 		}, nil
// 	})
// 	cfg.BaseEndpoint = aws.String(endpoint)
// 	fmt.Printf("AWS S3 options -> endpoint: %s, access key: %s, secret key: %s\n", endpoint, accessKey, secretKey)

// 	client := s3.NewFromConfig(cfg)

// 	return client
// }

func CreateAWSS3Client(endpoint, accessKey, secretKey, bucketName string) *s3.S3 {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession := session.New(s3Config)

	s3Client := s3.New(newSession)

	cparams := &s3.CreateBucketInput{
		Bucket: &bucketName,
	}

	_, err := s3Client.CreateBucket(cparams)
	if err != nil {
		fmt.Println(err.Error())
	}

	return s3Client
}
