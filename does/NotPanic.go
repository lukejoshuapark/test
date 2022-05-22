package does

import (
	"reflect"

	"github.com/lukejoshuapark/test"
)

// NotPanicAssertion provides the assertion logic for NotPanic.
type NotPanicAssertion struct{}

var _ test.Assertion = &NotPanicAssertion{}

// NotPanic calls the provided function and returns a passed result if the
// call does not panic.
var NotPanic = &NotPanicAssertion{}

// Evaluate evaluates the provided value in the context of the assertion.
func (a *NotPanicAssertion) Evaluate(actual interface{}) test.Result {
	xv := reflect.ValueOf(actual)
	xt := xv.Type()

	if xt.Kind() != reflect.Func {
		return test.Result{
			Pass:    false,
			Message: "expected a function",
			Specifics: map[string]interface{}{
				"Type": xt,
			},
		}
	}

	if xt.NumIn() > 0 || xt.NumOut() > 0 {
		return test.Result{
			Pass:    false,
			Message: "expected a function that takes no inputs and returns nothing",
			Specifics: map[string]interface{}{
				"Type": xt,
			},
		}
	}

	r := func() (r interface{}) {
		defer func() {
			r = recover()
		}()

		xv.Call([]reflect.Value{})
		return
	}()

	if r != nil {
		return test.Result{
			Pass:    false,
			Message: "expected function call to not panic",
			Specifics: map[string]interface{}{
				"Panic Result": r,
			},
		}
	}

	return test.PassedResult
}
