package is

import "testing"

func TestTrue(t *testing.T) {
	testCases := []struct {
		given      interface{}
		shouldPass bool
	}{
		{given: "Hello, World!", shouldPass: false},
		{given: 5, shouldPass: false},
		{given: false, shouldPass: false},
		{given: true, shouldPass: true},
	}

	for i, testCase := range testCases {
		// Arrange and Act.
		r := True.Evaluate(testCase.given)

		// Assert.
		if r.Pass != testCase.shouldPass {
			t.Fatalf("expected testCase %v Pass to be %v but was %v: %v", i, testCase.shouldPass, r.Pass, r.Message)
		}
	}
}
