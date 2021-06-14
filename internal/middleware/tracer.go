package middleware

import (
	"context"

	"github.com/uber/jaeger-client-go"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/tianmai777/blog/global"
)

func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx context.Context
		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(span.Context()))
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path)
		}
		defer span.Finish()
		c.Request = c.Request.WithContext(ctx)

		var (
			traceID     string
			spanID      string
			spanContext = span.Context()
		)
		switch spanContext.(type) {
		case jaeger.SpanContext:
			traceID = spanContext.(jaeger.SpanContext).TraceID().String()
			spanID = spanContext.(jaeger.SpanContext).SpanID().String()
		}
		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", spanID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
