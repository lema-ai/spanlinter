package other

import "context"

// StartSpan is an unrelated helper with the same name
func StartSpan(ctx context.Context, name string) context.Context {
	return ctx
}
