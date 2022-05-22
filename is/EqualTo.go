package is

import (
	"reflect"

	"gitlab.com/ljpcore/golib/test"
)

// EqualToAssertion provides the assertion logic for EqualTo.
type EqualToAssertion struct {
	expected interface{}
}

var _ test.Assertion = &EqualToAssertion{}

// EqualTo determines if one value is equal to another.  This assertion uses the
// standard == condition on two values of type interface{} when evaluating.
// You may prefer to use is.DeepEqualTo which evaluates using the deep-equal
// rules of the reflect package.
func EqualTo(expected interface{}) *EqualToAssertion {
	return &EqualToAssertion{expected: expected}
}

// Evaluate evaluates the provided value in the context of the assertion.
func (a *EqualToAssertion) Evaluate(actual interface{}) test.Result {
	if actual != a.expected {
		return test.Result{
			Pass:    false,
			Message: "expected equal values",
			Specifics: map[string]interface{}{
				"actual       ": actual,
				"actual Type  ": reflect.TypeOf(actual),
				"expected     ": a.expected,
				"expected Type": reflect.TypeOf(a.expected),
			},
		}
	}

	return test.PassedResult
}
