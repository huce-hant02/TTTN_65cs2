package logger

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"os"
	"strings"
)

const TraceKey = "trace_id"

func GenerateTraceID(serviceName string) string {
	uUid := uuid.NewString()
	traceID := fmt.Sprintf("%s-%s", serviceName, strings.ReplaceAll(uUid, "-", ""))
	return traceID
}

func NewContextWithTraceID(ctx context.Context, serviceName string) context.Context {
	return context.WithValue(ctx, TraceKey, GenerateTraceID(serviceName))
}

func NewBackgroundContextWithTraceID(serviceName string) context.Context {
	return NewContextWithTraceID(context.Background(), serviceName)
}

func NewLogger(ctx context.Context) *log.Helper {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
	// Set trace id for logger
	if ctx.Value(TraceKey) == nil {
		return log.NewHelper(logger).WithContext(ctx)
	}
	traceID := ctx.Value(TraceKey).(string)
	logger = log.With(logger, TraceKey, traceID)
	return log.NewHelper(logger).WithContext(ctx)
}
