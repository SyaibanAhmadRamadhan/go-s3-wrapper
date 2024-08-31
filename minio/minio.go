package s3_wrapper_minio

import (
	"github.com/minio/minio-go/v7"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type s3minio struct {
	client *minio.Client
	tracer trace.Tracer
	attrs  []attribute.KeyValue
}

func New(client *minio.Client) *s3minio {
	tp := otel.GetTracerProvider()
	return &s3minio{
		client: client,
		tracer: tp.Tracer("s3.wrapper.minio.tracer", trace.WithInstrumentationVersion("v1.0.0")),
		attrs: []attribute.KeyValue{
			attribute.String("minio.library.version", "v7"),
		},
	}
}
