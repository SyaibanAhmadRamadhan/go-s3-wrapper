package s3wrapper

import "context"

type S3Client interface {
	PresignedUrlUploadObject(ctx context.Context, input PresignedUrlUploadObjectInput) (output PresignedUrlUploadObjectOutput, err error)
	PresignedUrlGetObject(ctx context.Context, input PresignedUrlGetObjectInput) (output PresignedUrlGetObjectOutput, err error)
}
