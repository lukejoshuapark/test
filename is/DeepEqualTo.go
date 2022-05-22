package is

import (
	"reflect"

	"github.com/lukejoshuapark/test"
)

// DeepEqualToAssertion provides the assertion logic for DeepEqualTo.
type DeepEqualToAssertion struct {
	expected interface{}
}

var _ test.Assertion = &DeepEqualToAssertion{}

// DeepEqualTo determines if one value is deep-equal to another, according to
// the rules of the reflect package.
func DeepEqualTo(expected interface{}) *DeepEqualToAssertion {
	return &DeepEqualToAssertion{expected: expected}
}

// Evaluate evaluates the provided value in the context of the assertion.
func (a *DeepEqualToAssertion) Evaluate(actual interface{}) test.Result {
	if !reflect.DeepEqual(actual, a.expected) {
		return test.Result{
			Pass:    false,
			Message: "expected deep-equal values",
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
