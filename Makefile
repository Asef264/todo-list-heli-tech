.PHONY: run 

setup:
	go get github.com/aws/aws-sdk-go-v2/aws
	go get github.com/aws/aws-sdk-go-v2/config
	go get github.com/aws/aws-sdk-go-v2/service/s3
	go get github.com/golang-migrate/migrate/v4
	go get github.com/golang-migrate/migrate/v4/database/postgres
	go get github.com/lib/pq
	go get github.com/minio/minio-go/v7
	go get github.com/minio/minio-go/v7/pkg/credentials
	go get github.com/gin-gonic/gin
	go get github.com/DATA-DOG/go-sqlmock
	go get github.com/stretchr/testify/assert

run :
	go run cmd/main.go