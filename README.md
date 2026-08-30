# spanlinter

golangci-lint plugin that checks `log.StartSpan` span names in lemmata.

- names must be dotted snake_case with at least two segments, e.g. `assessment.generate_report`
- no two call sites may use the same name (constants counted per call site)
- non-constant names are skipped
