package b

import (
	"context"

	"github.com/lema.ai/lemmata/services/internal/log"
)

// Package b does not import package a. The duplicate must still be found, which
// is what the package-level map buys over an analysis.Fact.
func testCrossPackageDuplicate(ctx context.Context) {
	log.StartSpan(ctx, "a.generate_report") // want `duplicate log.StartSpan name "a.generate_report", also used at .*a\.go:`
}

func testUnique(ctx context.Context) {
	log.StartSpan(ctx, "b.unique_name")
}
