package filestorage

import (
	"bytes"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

const AclPublicRead = "public-read"

var ErrStorageServiceNotAvailable = errors.New("Error uploading token image, please try again later")

type Acl string

type Storage struct {
	awsKey    string
	awsSecret string
	token     string
	region    string
	bucket    string
}

func GetStorage(awsKey, awsSecret string) datatype.FileStorage {
	return &Storage{
		awsKey:    awsKey,
		awsSecret: awsSecret,
		token:     "",
		region:    "us-east-1",
		bucket:    "anychange",
	}
}

func (s *Storage) getS3() (*s3.S3, error) {
	creds := credentials.NewStaticCredentials(s.awsKey, s.awsSecret, s.token)
	_, err := creds.Get()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Storage:getAwsSession:!")
		return nil, ErrStorageServiceNotAvailable
	}
	cfg := aws.NewConfig().WithRegion(s.region).WithCredentials(creds)
	return s3.New(session.New(), cfg), nil
}

func (s *Storage) Put(
	path string,
	contentType string,
	content []byte,
	acl string,
) error {
	svc, err := s.getS3()
	if err != nil {
		return err
	}
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
		Body:   bytes.NewReader(content),
		//ContentLength: aws.Int64(size),
		ContentType: aws.String(contentType),
		ACL:         &acl,
	})
	return err
}
