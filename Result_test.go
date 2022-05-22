package test

import "testing"

func TestResultStringPass(t *testing.T) {
	// Arrange.
	r := PassedResult

	// Act.
	str := r.String()

	// Assert.
	if str != "" {
		t.Fatalf(`expected String to return an empty string, but got "%v"`, str)
	}
}

func TestResultStringFail(t *testing.T) {
	// Arrange.
	r := Result{
		Pass:    false,
		Message: "The test failed",
		Specifics: map[string]interface{}{
			"x": 123,
			"y": "abc",
		},
	}

	// Act.
	str := r.String()

	// Assert.
	expectedStr := "The test failed\nx: 123\ny: abc"
	if str != expectedStr {
		t.Fatalf("expected String to return:\n%v\nbut returned\n%v", expectedStr, str)
	}
}
