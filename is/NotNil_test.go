package is

import (
	"bytes"
	"testing"
)

func TestNotNil(t *testing.T) {
	testCases := []struct {
		given      interface{}
		shouldPass bool
	}{
		{given: "Hello, World!", shouldPass: true},
		{given: 5, shouldPass: true},
		{given: false, shouldPass: true},
		{given: true, shouldPass: true},
		{given: nil, shouldPass: false},
		{given: (*bytes.Buffer)(nil), shouldPass: false},
		{given: map[string]string(nil), shouldPass: false},
		{given: []byte(nil), shouldPass: false},
		{given: (func() bool)(nil), shouldPass: false},
		{given: chan int(nil), shouldPass: false},
	}

	for i, testCase := range testCases {
		// Arrange and Act.
		r := NotNil.Evaluate(testCase.given)

		// Assert.
		if r.Pass != testCase.shouldPass {
			t.Fatalf("expected testCase %v Pass to be %v but was %v: %v", i, testCase.shouldPass, r.Pass, r.Message)
		}
	}
}
