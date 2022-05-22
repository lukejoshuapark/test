package is

import (
	"reflect"

	"gitlab.com/ljpcore/golib/test"
)

// NilAssertion provides the assertion logic for Nil.
type NilAssertion struct{}

var _ test.Assertion = &NilAssertion{}

// Nil determines if a value is equal to nil.  An interface with a type but no
// value is considered nil.
var Nil = &NilAssertion{}

// Evaluate evaluates the provided value in the context of the assertion.
func (*NilAssertion) Evaluate(actual interface{}) test.Result {
	xt := reflect.TypeOf(actual)

	if actual == nil {
		return test.PassedResult
	}

	xv := reflect.ValueOf(actual)
	if isKindNilable(xv.Kind()) && xv.IsNil() {
		return test.PassedResult
	}

	return test.Result{
		Pass:    false,
		Message: "expected value to be nil",
		Specifics: map[string]interface{}{
			"actual": actual,
			"Type  ": xt,
		},
	}
}

func isKindNilable(kind reflect.Kind) bool {
	return kind == reflect.Ptr ||
		kind == reflect.Slice ||
		kind == reflect.Map ||
		kind == reflect.Func ||
		kind == reflect.Chan
}
