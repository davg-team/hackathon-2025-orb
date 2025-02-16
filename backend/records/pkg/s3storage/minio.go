package s3storage

import (
	"context"

	"github.com/davg/records/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStorage struct {
	client *minio.Client
}

var bucketName string

func Connect() *MinioStorage {
	cfg := config.Config().Bucket
	bucketName = cfg.BucketName
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}

	return &MinioStorage{client: minioClient}
}

func (s *MinioStorage) GetFile(ctx context.Context, objectKey string) (*minio.Object, error) {
	return s.client.GetObject(ctx, bucketName, objectKey, minio.GetObjectOptions{})
}
