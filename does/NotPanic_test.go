package does

import (
	"fmt"
	"testing"
)

func TestNotPanic(t *testing.T) {
	f1 := 5
	f2 := func(x int) {}
	f3 := func() bool { return false }
	f4 := func() { panic("ahh!") }
	f5 := func() {}

	testCases := []struct {
		given      interface{}
		shouldPass bool
	}{
		{given: f1, shouldPass: false},
		{given: f2, shouldPass: false},
		{given: f3, shouldPass: false},
		{given: f4, shouldPass: false},
		{given: f5, shouldPass: true},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			// Arrange and Act.
			r := NotPanic.Evaluate(testCase.given)

			// Assert.
			if r.Pass != testCase.shouldPass {
				t.Fatalf("expected Pass to be %v but was %v: %v", testCase.shouldPass, r.Pass, r.Message)
			}
		})
	}
}
