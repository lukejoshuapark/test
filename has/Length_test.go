package has

import "testing"

func TestLength(t *testing.T) {
	c1 := make(chan int, 2)
	c2 := make(chan int, 2)
	c1 <- 0
	c2 <- 0
	c2 <- 0

	testCases := []struct {
		given      interface{}
		shouldPass bool
	}{
		{given: "Hello, World!", shouldPass: false},
		{given: "He", shouldPass: true},
		{given: []byte{1, 2, 3}, shouldPass: false},
		{given: []byte{1, 2}, shouldPass: true},
		{given: [1]byte{}, shouldPass: false},
		{given: [2]byte{}, shouldPass: true},
		{given: map[int]int{0: 0}, shouldPass: false},
		{given: map[int]int{0: 0, 1: 1}, shouldPass: true},
		{given: c1, shouldPass: false},
		{given: c2, shouldPass: true},
		{given: 5, shouldPass: false},
	}

	for i, testCase := range testCases {
		// Arrange and Act.
		r := Length(2).Evaluate(testCase.given)

		// Assert.
		if r.Pass != testCase.shouldPass {
			t.Fatalf("expected testCase %v Pass to be %v but was %v: %v", i, testCase.shouldPass, r.Pass, r.Message)
		}
	}
}
