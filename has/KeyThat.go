package has

import (
	"reflect"

	"gitlab.com/ljpcore/golib/test"
)

// KeyThatAssertion provides the assertion logic for KeyThat.
type KeyThatAssertion struct {
	expected test.Assertion
}

var _ test.Assertion = &KeyThatAssertion{}

// KeyThat determines if the provided map has a key that passes the provided
// assertion.
func KeyThat(expected test.Assertion) *KeyThatAssertion {
	return &KeyThatAssertion{expected: expected}
}

// Evaluate evaluates the provided value in the context of the assertion.
func (a *KeyThatAssertion) Evaluate(actual interface{}) test.Result {
	xt := reflect.TypeOf(actual)
	if xt.Kind() != reflect.Map {
		return test.Result{
			Pass:    false,
			Message: "expected a map",
			Specifics: map[string]interface{}{
				"Type": xt,
			},
		}
	}

	xv := reflect.ValueOf(actual)
	xks := xv.MapKeys()

	if len(xks) == 0 {
		return test.Result{
			Pass:    false,
			Message: "map was nil or empty, so contained no keys",
			Specifics: map[string]interface{}{
				"Type": xt,
			},
		}
	}

	for _, kv := range xks {
		internalResult := a.expected.Evaluate(kv.Interface())
		if internalResult.Pass {
			return test.PassedResult
		}
	}

	return test.Result{
		Pass:    false,
		Message: "expected map to contain a key that passed assertion",
		Specifics: map[string]interface{}{
			"Length": xv.Len(),
			"Type  ": xt,
		},
	}
}
