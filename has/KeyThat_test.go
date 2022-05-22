package has

import (
	"testing"

	"gitlab.com/ljpcore/golib/test/is"
)

func TestKeyThat(t *testing.T) {
	testCases := []struct {
		given      interface{}
		shouldPass bool
	}{
		{given: "Hello, World!", shouldPass: false},
		{given: map[string]int{"b": 0}, shouldPass: false},
		{given: map[string]int{"a": 0}, shouldPass: true},
		{given: map[string]int{}, shouldPass: false},
	}

	for i, testCase := range testCases {
		// Arrange and Act.
		r := KeyThat(is.EqualTo("a")).Evaluate(testCase.given)

		// Assert.
		if r.Pass != testCase.shouldPass {
			t.Fatalf("expected testCase %v Pass to be %v but was %v: %v", i, testCase.shouldPass, r.Pass, r.Message)
		}
	}
}
