package documents

import (
	"github/joaltoroc/storicard/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewS3(cfg config.Config) (*s3.S3, error) {
	sessionData, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWS.Region),
	})
	if err != nil {
		return nil, err
	}

	return s3.New(sessionData), nil
}
