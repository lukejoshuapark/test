package is

import (
	"bytes"
	"fmt"
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
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			// Arrange and Act.
			r := NotNil.Evaluate(testCase.given)

			// Assert.
			if r.Pass != testCase.shouldPass {
				t.Fatalf("expected Pass to be %v but was %v: %v", testCase.shouldPass, r.Pass, r.Message)
			}
		})
	}
}
