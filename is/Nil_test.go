package is

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestNil(t *testing.T) {
	testCases := []struct {
		given      interface{}
		shouldPass bool
	}{
		{given: "Hello, World!", shouldPass: false},
		{given: 5, shouldPass: false},
		{given: false, shouldPass: false},
		{given: true, shouldPass: false},
		{given: nil, shouldPass: true},
		{given: (*bytes.Buffer)(nil), shouldPass: true},
		{given: map[string]string(nil), shouldPass: true},
		{given: []byte(nil), shouldPass: true},
		{given: (func() bool)(nil), shouldPass: true},
		{given: chan int(nil), shouldPass: true},
		{given: io.Writer((*os.File)(nil)), shouldPass: true},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			// Arrange and Act.
			r := Nil.Evaluate(testCase.given)

			// Assert.
			if r.Pass != testCase.shouldPass {
				t.Fatalf("expected Pass to be %v but was %v: %v", testCase.shouldPass, r.Pass, r.Message)
			}
		})
	}
}
