package a

import (
	"context"

	"github.com/lema.ai/lemmata/services/internal/log"
	"other"
)

const spanShared = "a.shared_name"

type service struct {
	name string
}

// Test case 1: a well-formed name is accepted
func testValid(ctx context.Context) {
	log.StartSpan(ctx, "a.generate_report")
}

// Test case 2: a constant with a well-formed name is accepted
func testValidConstant(ctx context.Context) {
	log.StartSpan(ctx, spanShared)
}

// Test case 3: PascalCase is rejected
func testPascalCase(ctx context.Context) {
	log.StartSpan(ctx, "GenerateReport") // want `span name "GenerateReport" must be dotted snake_case with at least two segments`
}

// Test case 4: a single segment is rejected
func testSingleSegment(ctx context.Context) {
	log.StartSpan(ctx, "generate_report") // want `span name "generate_report" must be dotted snake_case with at least two segments`
}

// Test case 5: a dotted name with a non-snake segment is rejected
func testMixedCaseSegment(ctx context.Context) {
	log.StartSpan(ctx, "a.generateReport") // want `span name "a.generateReport" must be dotted snake_case with at least two segments`
}

// Test case 6: a second call site using the same literal is a duplicate
func testDuplicateLiteral(ctx context.Context) {
	log.StartSpan(ctx, "a.generate_report") // want `duplicate log.StartSpan name "a.generate_report", also used at .*a\.go:`
}

// Test case 7: a second call site using the same constant is a duplicate
func testDuplicateConstant(ctx context.Context) {
	log.StartSpan(ctx, spanShared) // want `duplicate log.StartSpan name "a.shared_name", also used at .*a\.go:`
}

// Test case 8: a non-constant name is skipped
func testNonConstant(ctx context.Context, s *service) {
	log.StartSpan(ctx, s.name)
}

// Test case 9: a concatenation with a non-constant operand is skipped
func testConcatenation(ctx context.Context, s *service) {
	log.StartSpan(ctx, "kafka-notifier:"+s.name)
}

// Test case 10: an unrelated StartSpan is not flagged
func testOtherPackage(ctx context.Context) {
	other.StartSpan(ctx, "NotOurSpan")
}
