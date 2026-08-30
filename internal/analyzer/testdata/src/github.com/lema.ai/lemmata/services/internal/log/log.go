package log

import "context"

// Span is a minimal stub of trace.Span for testing
type Span struct{}

// StartSpan is a stub of the real helper
func StartSpan(ctx context.Context, name string) (context.Context, *Span) {
	return ctx, &Span{}
}
