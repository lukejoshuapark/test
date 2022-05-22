package has

import (
	"reflect"

	"gitlab.com/ljpcore/golib/test"
)

// ValueThatAssertion provides the assertion logic for ValueThat.
type ValueThatAssertion struct {
	expected test.Assertion
}

var _ test.Assertion = &ValueThatAssertion{}

// ValueThat determines if the provided array, map or slice has a value that
// passes the provided assertion.
func ValueThat(expected test.Assertion) *ValueThatAssertion {
	return &ValueThatAssertion{expected: expected}
}

// Evaluate evaluates the provided value in the context of the assertion.
func (a *ValueThatAssertion) Evaluate(actual interface{}) test.Result {
	xt := reflect.TypeOf(actual)
	xk := xt.Kind()

	if xk != reflect.Array && xk != reflect.Map && xk != reflect.Slice {
		return test.Result{
			Pass:    false,
			Message: "expected an array, map or slice",
			Specifics: map[string]interface{}{
				"Type": xt,
			},
		}
	}

	xv := reflect.ValueOf(actual)
	xvs := getValuesFor(xv)

	if len(xvs) == 0 {
		return test.Result{
			Pass:    false,
			Message: "set of values was nil or empty, so contained no values",
			Specifics: map[string]interface{}{
				"Type": xt,
			},
		}
	}

	for _, kv := range xvs {
		internalResult := a.expected.Evaluate(kv)
		if internalResult.Pass {
			return test.PassedResult
		}
	}

	return test.Result{
		Pass:    false,
		Message: "expected set of values to contain a value that passed assertion",
		Specifics: map[string]interface{}{
			"Length": xv.Len(),
			"Type  ": xt,
		},
	}
}

func getValuesFor(xv reflect.Value) []interface{} {
	if xv.Kind() == reflect.Map {
		return getValuesForMap(xv)
	}

	return getValuesForArrayOrSlice(xv)
}

func getValuesForMap(xv reflect.Value) []interface{} {
	r := make([]interface{}, 0, xv.Len())

	for _, k := range xv.MapKeys() {
		r = append(r, xv.MapIndex(k).Interface())
	}

	return r
}

func getValuesForArrayOrSlice(xv reflect.Value) []interface{} {
	r := make([]interface{}, xv.Len())

	for i := 0; i < xv.Len(); i++ {
		r[i] = xv.Index(i).Interface()
	}

	return r
}
