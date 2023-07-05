package s3

import (
	"bytes"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// s3API defines the methods used from the AWS SDK's S3 client.
type s3API interface {
	GetObject(*s3.GetObjectInput) (*s3.GetObjectOutput, error)
}

type Client struct {
	Bucket string
	S3     s3API
}

type Config struct {
	Bucket    string
	Region    string
	AccessKey string
	SecretKey string
	Endpoint  string
}

func NewClient(cfg *Config) *Client {
	awsSession := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String(cfg.Region),
		Credentials:      credentials.NewSharedCredentials("", "default"),
		Endpoint:         aws.String(cfg.Endpoint),
		S3ForcePathStyle: aws.Bool(true),
	}))

	s3Client := s3.New(awsSession)

	return &Client{
		Bucket: cfg.Bucket,
		S3:     s3Client,
	}
}

func (c *Client) Get(path string) ([]byte, error) {
	resp, err := c.S3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
