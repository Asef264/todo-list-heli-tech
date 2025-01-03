.PHONY: run 

setup:
	go get github.com/golang-migrate/migrate/v4
	go get github.com/golang-migrate/migrate/v4/database/postgres
	go get github.com/lib/pq
	go get github.com/minio/minio-go/v7
	go get github.com/gin-gonic/gin
	go get github.com/DATA-DOG/go-sqlmock
	go get github.com/stretchr/testify/assert
	go get github.com/minio/minio-go/v7/pkg/credentials
	go get github.com/aws/aws-sdk-go/aws
	go get github.com/aws/aws-sdk-go/service/s3

run :
	go run cmd/main.go


unit_test:
	go test -v ./internal/ports

bench_mark:
	go test -bench=. ./internal/ports

docker_compose_run:
	docker compose up