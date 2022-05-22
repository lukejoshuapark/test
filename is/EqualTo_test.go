package is

import (
	"io"
	"os"
	"testing"
)

func TestEqualTo(t *testing.T) {
	testCases := []struct {
		actual     interface{}
		expected   interface{}
		shouldPass bool
	}{
		{actual: 4, expected: 5, shouldPass: false},
		{actual: int16(5), expected: 5, shouldPass: false},
		{actual: 5, expected: 5, shouldPass: true},
		{actual: "Hello, World", expected: "Hello, World!", shouldPass: false},
		{actual: "Hello, World!", expected: "Hello, World!", shouldPass: true},
		{actual: nil, expected: nil, shouldPass: true},
		{actual: [2]byte{}, expected: [2]byte{}, shouldPass: true},
		{actual: io.Writer((*os.File)(nil)), expected: nil, shouldPass: false},
	}

	for i, testCase := range testCases {
		// Arrange and Act.
		r := EqualTo(testCase.expected).Evaluate(testCase.actual)

		// Assert.
		if r.Pass != testCase.shouldPass {
			t.Fatalf("expected testCase %v Pass to be %v but was %v: %v", i, testCase.shouldPass, r.Pass, r.Message)
		}
	}
}
