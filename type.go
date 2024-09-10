package s3wrapper

import (
	"io"
	"time"
)

type PresignedUrlUploadObjectInput struct {
	BucketName string
	Path       string
	MimeType   string
	Checksum   string
	Expired    time.Duration
}

type PresignedUrlUploadObjectOutput struct {
	URL       string
	ExpiredAt time.Time

	// MinioFormData if you can use minio s3
	MinioFormData map[string]string
}

type PresignedUrlGetObjectInput struct {
	ObjectName string
	BucketName string
	Expired    time.Duration
}

type PresignedUrlGetObjectOutput struct {
	URL string
}

type GetObjectInput struct {
	ObjectName string
	BucketName string
}

type GetObjectOutput struct {
	Object io.ReadCloser
}
