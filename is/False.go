package is

import (
	"fmt"
	"reflect"

	"github.com/lukejoshuapark/test"
)

// FalseAssertion provides the assertion logic for False.
type FalseAssertion struct{}

var _ test.Assertion = &FalseAssertion{}

// False determines if a value is the boolean value false.
var False = &FalseAssertion{}

// Evaluate evaluates the provided value in the context of the assertion.
func (*FalseAssertion) Evaluate(actual interface{}) test.Result {
	et := reflect.TypeOf(false)
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

	if actual.(bool) {
		return test.Result{
			Pass:    false,
			Message: "expected the expression to be false, but was true",
			Specifics: map[string]interface{}{
				"actual": actual,
			},
		}
	}

	return test.PassedResult
}
