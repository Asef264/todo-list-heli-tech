package s3Adaptor

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	s3Config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

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

func CreateMinioClient(endpoint string, accessKeyID string,
	secretAccessKey string, useSSL bool,
) (*minio.Client, error) {
	if endpoint == "" {
		return nil, fmt.Errorf("endpoint must not be empty")
	}

	tlsConfig := &tls.Config{}
	if useSSL {
		tlsConfig.InsecureSkipVerify = true
	}

	var transport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
		DisableCompression:    true,
	}

	mClient, err := minio.New(endpoint, &minio.Options{
		Creds:     credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure:    useSSL,
		Transport: transport,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	return mClient, nil
}
