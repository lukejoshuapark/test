package test

import "fmt"

// That asserts that the provided value passes the assertion.  It will fail t
// if it does not.  Any suffixed values will be included in the output should
// the test fail.
func That(t TLike, actual interface{}, assertion Assertion, args ...interface{}) {
	t.Helper()

	if r := assertion.Evaluate(actual); !r.Pass {
		t.Fatalf("%v", formatFatalMessage(r, args...))
	}
}

func formatFatalMessage(r Result, args ...interface{}) string {
	argStr := ""
	for _, arg := range args {
		argStr += fmt.Sprintf("%v\n", arg)
	}

	if argStr != "" {
		argStr += "\n"
	}

	return fmt.Sprintf("%v%v", argStr, r)
}
