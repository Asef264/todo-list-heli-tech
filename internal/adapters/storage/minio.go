package storage

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

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
