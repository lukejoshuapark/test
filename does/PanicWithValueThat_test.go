package does

import (
	"testing"

	"gitlab.com/ljpcore/golib/test/is"
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
		// Arrange and Act.
		r := PanicWithValueThat(is.True).Evaluate(testCase.given)

		// Assert.
		if r.Pass != testCase.shouldPass {
			t.Fatalf("expected testCase %v Pass to be %v but was %v: %v", i, testCase.shouldPass, r.Pass, r.Message)
		}
	}
}