package has

import (
	"fmt"
	"reflect"

	"gitlab.com/ljpcore/golib/test"
)

// LengthAssertion provides the assertion logic for Length.
type LengthAssertion struct {
	expected int
}

var _ test.Assertion = &LengthAssertion{}

// Length determines if the provided array, channel, map, slice or string has
// a given length.
func Length(expected int) *LengthAssertion {
	return &LengthAssertion{expected: expected}
}

// Evaluate evaluates the provided value in the context of the assertion.
func (a *LengthAssertion) Evaluate(actual interface{}) test.Result {
	allowedKinds := map[reflect.Kind]struct{}{
		reflect.Array:  {},
		reflect.Chan:   {},
		reflect.Map:    {},
		reflect.Slice:  {},
		reflect.String: {},
	}

	xv := reflect.ValueOf(actual)
	xt := xv.Type()
	xk := xt.Kind()

	if _, ok := allowedKinds[xk]; !ok {
		return test.Result{
			Pass:    false,
			Message: "expected a type that has a length",
			Specifics: map[string]interface{}{
				"Type": xt,
			},
		}
	}

	xl := xv.Len()
	if xl != a.expected {
		return test.Result{
			Pass:    false,
			Message: fmt.Sprintf("expected a length of %v, but was %v", a.expected, xl),
			Specifics: map[string]interface{}{
				"Actual Length  ": xl,
				"Expected Length": a.expected,
				"Type           ": xt,
			},
		}
	}

	return test.PassedResult
}
