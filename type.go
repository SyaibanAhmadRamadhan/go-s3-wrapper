package s3wrapper

import "time"

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
