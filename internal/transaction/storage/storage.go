package storage

import (
	"github/joaltoroc/storicard/config"
	"github/joaltoroc/storicard/internal/transaction"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type storage struct {
	s3  *s3.S3
	cfg config.Config
}

func NewStorage(s3 *s3.S3, cfg config.Config) transaction.Storage {
	return &storage{
		s3,
		cfg,
	}
}

func (s *storage) GetFile(fileName string) ([]byte, error) {
	request := &s3.GetObjectInput{
		Bucket: aws.String(s.cfg.AWS.S3.Bucket),
		Key:    aws.String(fileName),
	}

	result, err := s.s3.GetObject(request)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
