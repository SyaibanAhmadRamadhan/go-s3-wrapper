package s3_wrapper_minio

import (
	"context"
	"errors"
	s3wrapper "github.com/SyaibanAhmadRamadhan/go-s3-wrapper"
	"github.com/minio/minio-go/v7"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (s *s3minio) GetObject(ctx context.Context, input s3wrapper.GetObjectInput) (output s3wrapper.GetObjectOutput, err error) {
	ctx, span := s.tracer.Start(ctx, "minio s3: Presigned Url Get Object", trace.WithAttributes(
		attribute.String("s3.minio.bucket.name", input.BucketName),
		attribute.String("s3.minio.object.name", input.ObjectName),
	))
	defer span.End()

	_, err = s.client.StatObject(ctx, input.BucketName, input.ObjectName, minio.StatObjectOptions{})
	if err != nil {
		minioErr := minio.ToErrorResponse(err)
		if minioErr.Code == "NoSuchKey" {
			err = errors.Join(err, s3wrapper.ErrObjectNotFound)
		} else {
			recordErrorOtel(span, err)
		}
		return output, errRecord(err)
	}

	object, err := s.client.GetObject(ctx, input.BucketName, input.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		recordErrorOtel(span, err)
		return output, errRecord(err)
	}

	output = s3wrapper.GetObjectOutput{
		Object: object,
	}
	return
}
