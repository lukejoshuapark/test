package is

import (
	"reflect"

	"github.com/lukejoshuapark/test"
)

// NotDeepEqualToAssertion provides the assertion logic for NotDeepEqualTo.
type NotDeepEqualToAssertion struct {
	expected interface{}
}

var _ test.Assertion = &NotDeepEqualToAssertion{}

// NotDeepEqualTo determines if one value is not deep-equal to another,
// according to the rules of the reflect package.
func NotDeepEqualTo(expected interface{}) *NotDeepEqualToAssertion {
	return &NotDeepEqualToAssertion{expected: expected}
}

// Evaluate evaluates the provided value in the context of the assertion.
func (a *NotDeepEqualToAssertion) Evaluate(actual interface{}) test.Result {
	if reflect.DeepEqual(actual, a.expected) {
		return test.Result{
			Pass:    false,
			Message: "expected not deep-equal values",
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
