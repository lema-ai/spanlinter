package a

// Deliberately empty. Its presence makes go/packages build both the "a" and
// "a [a.test]" variants, so every call site in a.go is visited twice — which is
// what exercises the same-position guard in the duplicate check.
