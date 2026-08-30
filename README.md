# spanlinter

golangci-lint plugin that checks `log.StartSpan` span names in lemmata.

- names must be dotted snake_case with at least two segments, e.g. `assessment.generate_report`
- no two call sites may use the same name (constants counted per call site)
- non-constant names are skipped

The duplicate check accumulates in a process-level map, so it only holds when the
whole tree is linted by one `golangci-lint run` with a cold cache. Splitting the
package patterns across separate invocations, or reusing a warm cache, hides
duplicates between packages that were not analyzed in the same process.
