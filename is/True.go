package is

import (
	"fmt"
	"reflect"

	"gitlab.com/ljpcore/golib/test"
)

// TrueAssertion provides the assertion logic for True.
type TrueAssertion struct{}

var _ test.Assertion = &TrueAssertion{}

// True determines if a value is the boolean value true.
var True = &TrueAssertion{}

// Evaluate evaluates the provided value in the context of the assertion.
func (*TrueAssertion) Evaluate(actual interface{}) test.Result {
	et := reflect.TypeOf(true)
	xt := reflect.TypeOf(actual)

	if xt != et {
		return test.Result{
			Pass:    false,
			Message: fmt.Sprintf("expected a bool value, but was provided a %v", xt),
			Specifics: map[string]interface{}{
				"Expected Type": et,
				"Given Type   ": xt,
				"actual       ": actual,
			},
		}
	}

	if !actual.(bool) {
		return test.Result{
			Pass:    false,
			Message: "expected the expression to be true, but was false",
			Specifics: map[string]interface{}{
				"actual": actual,
			},
		}
	}

	return test.PassedResult
}
