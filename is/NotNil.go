package is

import (
	"reflect"

	"github.com/lukejoshuapark/test"
)

// NotNilAssertion provides the assertion logic for NotNil.
type NotNilAssertion struct{}

var _ test.Assertion = &NotNilAssertion{}

// NotNil determines if a value is not equal to nil.  An interface with a type
// but no value is considered nil.
var NotNil = &NotNilAssertion{}

// Evaluate evaluates the provided value in the context of the assertion.
func (*NotNilAssertion) Evaluate(actual interface{}) test.Result {
	xt := reflect.TypeOf(actual)

	if actual != nil {
		xv := reflect.ValueOf(actual)
		if !isKindNilable(xv.Kind()) || !xv.IsNil() {
			return test.PassedResult
		}
	}

	return test.Result{
		Pass:    false,
		Message: "expected value to not be nil",
		Specifics: map[string]interface{}{
			"Type  ": xt,
			"actual": actual,
		},
	}
}
