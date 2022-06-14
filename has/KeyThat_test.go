package has

import (
	"fmt"
	"testing"

	"github.com/lukejoshuapark/test/is"
)

func TestKeyThat(t *testing.T) {
	testCases := []struct {
		given      interface{}
		shouldPass bool
	}{
		{given: "Hello, World!", shouldPass: false},
		{given: map[string]int{"b": 0}, shouldPass: false},
		{given: map[string]int{"a": 0}, shouldPass: true},
		{given: map[string]int{}, shouldPass: false},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			// Arrange and Act.
			r := KeyThat(is.EqualTo("a")).Evaluate(testCase.given)

			// Assert.
			if r.Pass != testCase.shouldPass {
				t.Fatalf("expected Pass to be %v but was %v: %v", testCase.shouldPass, r.Pass, r.Message)
			}
		})
	}
}
