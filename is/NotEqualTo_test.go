package is

import "testing"

func TestNotEqualTo(t *testing.T) {
	testCases := []struct {
		actual     interface{}
		expected   interface{}
		shouldPass bool
	}{
		{actual: 4, expected: 5, shouldPass: true},
		{actual: int16(5), expected: 5, shouldPass: true},
		{actual: 5, expected: 5, shouldPass: false},
		{actual: "Hello, World", expected: "Hello, World!", shouldPass: true},
		{actual: "Hello, World!", expected: "Hello, World!", shouldPass: false},
		{actual: nil, expected: nil, shouldPass: false},
		{actual: [2]byte{}, expected: [2]byte{}, shouldPass: false},
	}

	for i, testCase := range testCases {
		// Arrange and Act.
		r := NotEqualTo(testCase.expected).Evaluate(testCase.actual)

		// Assert.
		if r.Pass != testCase.shouldPass {
			t.Fatalf("expected testCase %v Pass to be %v but was %v: %v", i, testCase.shouldPass, r.Pass, r.Message)
		}
	}
}
