package adapters

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

// Mock S3 client
type mockS3Client struct {
	PutObjectOutput *s3.PutObjectOutput
	GetObjectOutput *s3.GetObjectOutput
	PutObjectError  error
	GetObjectError  error
	MockStorage     map[string][]byte
}

func (m *mockS3Client) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	if m.PutObjectError != nil {
		return nil, m.PutObjectError
	}
	// reading the content from the input.Body (io.ReadSeeker)
	var buf bytes.Buffer
	_, err := io.Copy(&buf, input.Body)
	if err != nil {
		return nil, err
	}
	m.MockStorage[*input.Key] = buf.Bytes()
	return m.PutObjectOutput, nil
}

func (m *mockS3Client) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	if m.GetObjectError != nil {
		return nil, m.GetObjectError
	}

	data, ok := m.MockStorage[*input.Key]
	if !ok {
		return nil, errors.New("not found")
	}

	return &s3.GetObjectOutput{
		Body: io.NopCloser(bytes.NewReader(data)),
	}, nil
}

func TestMockPutObject(t *testing.T) {
	mockClient := &mockS3Client{
		PutObjectOutput: &s3.PutObjectOutput{},
		MockStorage:     make(map[string][]byte),
	}

	input := &s3.PutObjectInput{
		Bucket: aws.String("test-bucket"),
		Key:    aws.String("test-key"),
		Body:   bytes.NewReader([]byte("test-content")),
	}

	_, err := mockClient.PutObject(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(mockClient.MockStorage["test-key"]) != "test-content" {
		t.Errorf("expected 'test-content', got '%s'", mockClient.MockStorage["test-key"])
	}
}

func BenchmarkMockPutObject(b *testing.B) {
	mockClient := &mockS3Client{
		PutObjectOutput: &s3.PutObjectOutput{},
		MockStorage:     make(map[string][]byte),
	}

	bucketName := "benchmark-bucket"
	keyName := "benchmark-key"
	content := []byte("benchmark content")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		input := &s3.PutObjectInput{
			Bucket: &bucketName,
			Key:    &keyName,
			Body:   bytes.NewReader(content),
		}

		_, err := mockClient.PutObject(input)
		assert.NoError(b, err, "PutObject should not return an error")
	}
}
