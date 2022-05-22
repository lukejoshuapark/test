package does

import (
	"fmt"
	"reflect"

	"gitlab.com/ljpcore/golib/test"
)

// PanicWithValueThatAssertion provides the assertion logic for
// PanicWithValueThat.
type PanicWithValueThatAssertion struct {
	expected test.Assertion
}

var _ test.Assertion = &PanicWithValueThatAssertion{}

// PanicWithValueThat calls the provided function and returns a passed result if
// the call panics and the panic value passes the internal assertion.
func PanicWithValueThat(expected test.Assertion) *PanicWithValueThatAssertion {
	return &PanicWithValueThatAssertion{
		expected: expected,
	}
}

// Evaluate evaluates the provided value in the context of the assertion.
func (a *PanicWithValueThatAssertion) Evaluate(actual interface{}) test.Result {
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

	if r == nil {
		return test.Result{
			Pass:    false,
			Message: "expected function call to panic",
		}
	}

	internalResult := a.expected.Evaluate(r)
	if internalResult.Pass {
		return test.PassedResult
	}

	return test.Result{
		Pass:      false,
		Message:   fmt.Sprintf("function call did panic, but panic value did not pass assertion: %v", internalResult.Message),
		Specifics: internalResult.Specifics,
	}
}
