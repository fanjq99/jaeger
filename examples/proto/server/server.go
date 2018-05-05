package server

import (
	"context"
	"sync"
	"time"

	pbJaeger "github.com/jaegertracing/jaeger/model/proto"
)

// Backend implements QueryServiceV1
type Backend struct {
	mu     sync.RWMutex
	traces []*pbJaeger.Trace
}

var _ pbJaeger.QueryServiceV1Server = (*Backend)(nil)

// New does new
func New() *Backend {
	return &Backend{}
}

// GetTrace gets trace
func (b *Backend) GetTrace(ctx context.Context, traceID *pbJaeger.TraceID) (*pbJaeger.Trace, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return &pbJaeger.Trace{
		Spans: []pbJaeger.Span{
			pbJaeger.Span{
				TraceID:       pbJaeger.TraceID{Low: 123},
				SpanID:        pbJaeger.SpanID{Value: 456},
				OperationName: "foo bar",
				StartTime:     time.Now(),
			},
		},
	}, nil
}
