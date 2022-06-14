package is

import (
	"fmt"
	"testing"
)

func TestNotDeepEqualTo(t *testing.T) {
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
		{actual: [2]byte{1, 1}, expected: [2]byte{1, 1}, shouldPass: false},
		{actual: []byte{1, 1}, expected: []byte{1, 1}, shouldPass: false},
		{actual: []byte(nil), expected: nil, shouldPass: true},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			// Arrange and Act.
			r := NotDeepEqualTo(testCase.expected).Evaluate(testCase.actual)

			// Assert.
			if r.Pass != testCase.shouldPass {
				t.Fatalf("expected Pass to be %v but was %v: %v", testCase.shouldPass, r.Pass, r.Message)
			}
		})
	}
}
