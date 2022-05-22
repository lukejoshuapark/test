package test

import (
	"fmt"
	"sort"
)

// Result is the result of evaluating a value using an assertion.  It is
// returned by the Evaluate method of types that implement Assertion.
//
// It specifies whether the value passed the assertion, and if it did not, why.
type Result struct {
	Pass      bool
	Message   string
	Specifics map[string]interface{}
}

var _ fmt.Stringer = &Result{}

// PassedResult is a Result with Pass set to true.
var PassedResult = Result{Pass: true}

// String formats the Result depending on the value of Pass.  If true, String
// returns an empty string.  If false, Message and Specifics are combined to
// provide an explanation for failure.
func (r Result) String() string {
	if r.Pass {
		return ""
	}

	keys := make([]string, len(r.Specifics))
	i := 0
	for k := range r.Specifics {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	str := r.Message
	for _, k := range keys {
		str += fmt.Sprintf("\n%v: %v", k, r.Specifics[k])
	}

	return str
}
