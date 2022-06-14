package does

import (
	"fmt"
	"testing"

	"github.com/lukejoshuapark/test/is"
)

func TestPanicWithValueThat(t *testing.T) {
	f1 := 5
	f2 := func(x int) {}
	f3 := func() bool { return false }
	f4 := func() { panic("ahh!") }
	f5 := func() {}
	f6 := func() { panic(true) }

	testCases := []struct {
		given      interface{}
		shouldPass bool
	}{
		{given: f1, shouldPass: false},
		{given: f2, shouldPass: false},
		{given: f3, shouldPass: false},
		{given: f4, shouldPass: false},
		{given: f5, shouldPass: false},
		{given: f6, shouldPass: true},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			// Arrange and Act.
			r := PanicWithValueThat(is.True).Evaluate(testCase.given)

			// Assert.
			if r.Pass != testCase.shouldPass {
				t.Fatalf("expected Pass to be %v but was %v: %v", testCase.shouldPass, r.Pass, r.Message)
			}
		})
	}
}
