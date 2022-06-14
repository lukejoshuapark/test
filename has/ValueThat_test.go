package has

import (
	"fmt"
	"testing"

	"github.com/lukejoshuapark/test/is"
)

func TestValueThat(t *testing.T) {
	testCases := []struct {
		given      interface{}
		shouldPass bool
	}{
		{given: "Hello, World!", shouldPass: false},
		{given: map[string]int{"b": 0}, shouldPass: false},
		{given: map[string]int{"a": 0, "b": 1}, shouldPass: false},
		{given: map[string]int{"a": 0, "b": 1, "c": 2}, shouldPass: true},
		{given: []int{0, 1}, shouldPass: false},
		{given: []int{0, 1, 2}, shouldPass: true},
		{given: [2]int{1, 3}, shouldPass: false},
		{given: [2]int{2, 1}, shouldPass: true},
		{given: []int{}, shouldPass: false},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			// Arrange and Act.
			r := ValueThat(is.EqualTo(2)).Evaluate(testCase.given)

			// Assert.
			if r.Pass != testCase.shouldPass {
				t.Fatalf("expected Pass to be %v but was %v: %v", testCase.shouldPass, r.Pass, r.Message)
			}
		})
	}
}
