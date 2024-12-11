package middlewares

import (
	"context"
	"sync"
)

// var sharable resources
var (
	logNo int = 1
	mu    sync.Mutex
)

func newTraceID() int {
	no := 1 // trace ID attached to HTTP request

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}

func SetTraceID(ctx context.Context, traceID int) context.Context {
	// add context to ctx {key, value}
	return context.WithValue(ctx, "traceID", traceID)
}

func GetTraceID(ctx context.Context) int {
	id := ctx.Value("traceID")

	// Value method returns any type, so type-assertion is needed
	if idInt, ok := id.(int); ok {
		return idInt
	}

	return 0
}
