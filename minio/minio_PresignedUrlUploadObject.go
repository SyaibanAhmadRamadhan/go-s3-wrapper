package s3_wrapper_minio

import (
	"context"
	s3wrapper "github.com/SyaibanAhmadRamadhan/go-s3-wrapper"
	"github.com/minio/minio-go/v7"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

func (s *s3minio) PresignedUrlUploadObject(ctx context.Context, input s3wrapper.PresignedUrlUploadObjectInput) (output s3wrapper.PresignedUrlUploadObjectOutput, err error) {
	expiredAt := time.Now().UTC().Add(input.Expired)

	ctx, span := s.tracer.Start(ctx, "minio s3 - Presigned Url upload object", trace.WithAttributes(
		attribute.String("s3.wrapper.minio.bucket.name", input.BucketName),
		attribute.String("s3.wrapper.minio.object.name", input.Path),
		attribute.String("s3.wrapper.minio.presigned.expired", expiredAt.Format(time.DateTime)),
		attribute.String("s3.wrapper.minio.object.mime_type", input.MimeType),
		attribute.String("s3.wrapper.minio.object.checksum_sha256", input.Checksum),
	))
	defer span.End()

	policy := minio.NewPostPolicy()
	policy.SetChecksum(minio.NewChecksum(minio.ChecksumSHA256, []byte(input.Checksum)))

	err = policy.SetKey(input.Path)
	if err != nil {
		recordErrorOtel(span, err)
		return output, errRecord(err)
	}

	err = policy.SetExpires(expiredAt)
	if err != nil {
		recordErrorOtel(span, err)
		return output, errRecord(err)
	}

	err = policy.SetContentLengthRange(1024, 2048)
	if err != nil {
		recordErrorOtel(span, err)
		return output, errRecord(err)
	}

	err = policy.SetBucket(input.BucketName)
	if err != nil {
		recordErrorOtel(span, err)
		return output, errRecord(err)
	}

	outputPresignedPostPolicy, formData, err := s.client.PresignedPostPolicy(ctx, policy)
	if err != nil {
		recordErrorOtel(span, err)
		return output, errRecord(err)
	}

	output = s3wrapper.PresignedUrlUploadObjectOutput{
		URL:           outputPresignedPostPolicy.String(),
		ExpiredAt:     expiredAt,
		MinioFormData: formData,
	}
	return
}
