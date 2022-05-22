package has

import (
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	typeStr := reflect.TypeOf("")
	typeInt := reflect.TypeOf(5)
	typeInt64 := reflect.TypeOf(int64(5))
	typeMap := reflect.TypeOf(map[string]string{})

	testCases := []struct {
		given      interface{}
		t          reflect.Type
		shouldPass bool
	}{
		{given: "Hello, World!", t: typeStr, shouldPass: true},
		{given: "Hello, World!", t: typeInt, shouldPass: false},
		{given: "Hello, World!", t: typeMap, shouldPass: false},
		{given: 5, t: typeInt, shouldPass: true},
		{given: 5, t: typeInt64, shouldPass: false},
		{given: int64(5), t: typeInt, shouldPass: false},
		{given: int64(5), t: typeInt64, shouldPass: true},
		{given: map[string]int{}, t: typeMap, shouldPass: false},
		{given: map[string]string{}, t: typeMap, shouldPass: true},
	}

	for i, testCase := range testCases {
		// Arrange and Act.
		r := Type(testCase.t).Evaluate(testCase.given)

		// Assert.
		if r.Pass != testCase.shouldPass {
			t.Fatalf("expected testCase %v Pass to be %v but was %v: %v", i, testCase.shouldPass, r.Pass, r.Message)
		}
	}
}
