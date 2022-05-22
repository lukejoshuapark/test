package has

import (
	"fmt"
	"reflect"

	"gitlab.com/ljpcore/golib/test"
)

// TypeAssertion provides the assertion logic for Type.
type TypeAssertion struct {
	expected reflect.Type
}

var _ test.Assertion = &TypeAssertion{}

// Type determines if the provided value has the type provided.
func Type(expected reflect.Type) *TypeAssertion {
	return &TypeAssertion{expected: expected}
}

// Evaluate evaluates the provided value in the context of the assertion.
func (a *TypeAssertion) Evaluate(actual interface{}) test.Result {
	xt := reflect.TypeOf(actual)
	if xt == a.expected {
		return test.PassedResult
	}

	return test.Result{
		Pass:    false,
		Message: fmt.Sprintf("expected value to have type %v, but had type %v", a.expected, xt),
		Specifics: map[string]interface{}{
			"Actual Type  ": xt,
			"Expected Type": a.expected,
		},
	}
}
