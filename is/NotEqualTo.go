package is

import (
	"reflect"

	"gitlab.com/ljpcore/golib/test"
)

// NotEqualToAssertion provides the assertion logic for NotEqualTo.
type NotEqualToAssertion struct {
	expected interface{}
}

var _ test.Assertion = &NotEqualToAssertion{}

// NotEqualTo determines if one value is not equal to another.  This assertion
// uses the standard != condition on two values of type interface{} when
// evaluating. You may prefer to use is.NotDeepEqualTo which evaluates using the
// deep-equal rules of the reflect package.
func NotEqualTo(expected interface{}) *NotEqualToAssertion {
	return &NotEqualToAssertion{expected: expected}
}

// Evaluate evaluates the provided value in the context of the assertion.
func (a *NotEqualToAssertion) Evaluate(x interface{}) test.Result {
	if x == a.expected {
		return test.Result{
			Pass:    false,
			Message: "expected not equal values",
			Specifics: map[string]interface{}{
				"actual       ": x,
				"actual Type  ": reflect.TypeOf(x),
				"expected     ": a.expected,
				"expected Type": reflect.TypeOf(a.expected),
			},
		}
	}

	return test.PassedResult
}
