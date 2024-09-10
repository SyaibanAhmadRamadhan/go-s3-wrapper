package s3_wrapper_minio

import (
	"context"
	s3wrapper "github.com/SyaibanAhmadRamadhan/go-s3-wrapper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/url"
)

func (s *s3minio) PresignedUrlGetObject(ctx context.Context, input s3wrapper.PresignedUrlGetObjectInput) (output s3wrapper.PresignedUrlGetObjectOutput, err error) {
	ctx, span := s.tracer.Start(ctx, "minio s3: Presigned Url Get Object", trace.WithAttributes(
		attribute.String("s3.minio.bucket.name", input.BucketName),
		attribute.String("s3.minio.object.name", input.ObjectName),
	))
	defer span.End()

	params := make(url.Values)

	getPresignedOutput, err := s.client.PresignedGetObject(ctx, input.BucketName, input.ObjectName, input.Expired, params)
	if err != nil {
		recordErrorOtel(span, err)
		return output, errRecord(err)
	}

	output = s3wrapper.PresignedUrlGetObjectOutput{
		URL: getPresignedOutput.String(),
	}
	return
}
