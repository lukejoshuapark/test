package is

import (
	"fmt"
	"testing"
)

func TestFalse(t *testing.T) {
	testCases := []struct {
		given      interface{}
		shouldPass bool
	}{
		{given: "Hello, World!", shouldPass: false},
		{given: 5, shouldPass: false},
		{given: false, shouldPass: true},
		{given: true, shouldPass: false},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			// Arrange and Act.
			r := False.Evaluate(testCase.given)

			// Assert.
			if r.Pass != testCase.shouldPass {
				t.Fatalf("expected Pass to be %v but was %v: %v", testCase.shouldPass, r.Pass, r.Message)
			}
		})
	}
}
