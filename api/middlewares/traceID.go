package middlewares

import "sync"

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
