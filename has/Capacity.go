package has

import (
	"fmt"
	"reflect"

	"github.com/lukejoshuapark/test"
)

// CapacityAssertion provides the assertion logic for Capacity.
type CapacityAssertion struct {
	expected int
}

var _ test.Assertion = &CapacityAssertion{}

// Capacity determines if the provided array, channel, map, slice or string has
// a given capacity.
func Capacity(expected int) *CapacityAssertion {
	return &CapacityAssertion{expected: expected}
}

// Evaluate evaluates the provided value in the context of the assertion.
func (a *CapacityAssertion) Evaluate(actual interface{}) test.Result {
	allowedKinds := map[reflect.Kind]struct{}{
		reflect.Array: {},
		reflect.Chan:  {},
		reflect.Slice: {},
	}

	xv := reflect.ValueOf(actual)
	xt := xv.Type()
	xk := xt.Kind()

	if _, ok := allowedKinds[xk]; !ok {
		return test.Result{
			Pass:    false,
			Message: "expected a type that has a capacity",
			Specifics: map[string]interface{}{
				"Type": xt,
			},
		}
	}

	xc := xv.Cap()
	if xc != a.expected {
		return test.Result{
			Pass:    false,
			Message: fmt.Sprintf("expected a capacity of %v, but was %v", a.expected, xc),
			Specifics: map[string]interface{}{
				"Actual Capacity  ": xc,
				"Expected Capacity": a.expected,
				"Type             ": xt,
			},
		}
	}

	return test.PassedResult
}
