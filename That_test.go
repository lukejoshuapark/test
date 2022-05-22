package test

import "testing"

func TestThatPass(t *testing.T) {
	// Arrange.
	mock := NewTMock()

	// Act.
	That(mock, 5, &equalToFive{})

	// Assert.
	if mock.HelperCallCount != 1 {
		t.Fatalf("expected a HelperCallCount of 1, but was %v", mock.HelperCallCount)
	}

	if mock.FatalMessage != "" {
		t.Fatalf(`expected FatalMessage to be an empty string, but was "%v"`, mock.FatalMessage)
	}
}

func TestThatFail(t *testing.T) {
	// Arrange.
	mock := NewTMock()

	// Act.
	That(mock, 4, &equalToFive{})

	// Assert.
	if mock.HelperCallCount != 1 {
		t.Fatalf("expected a HelperCallCount of 1, but was %v", mock.HelperCallCount)
	}

	expectedFatalMessage := "The provided value was not equal to 5\nExpected: 5\nGot: 4"
	if mock.FatalMessage != expectedFatalMessage {
		t.Fatalf("expected a FatalMessage of:\n%v\nbut returned\n%v", expectedFatalMessage, mock.FatalMessage)
	}
}

func TestThatFailWithArgs(t *testing.T) {
	// Arrange.
	mock := NewTMock()

	// Act.
	That(mock, 4, &equalToFive{}, "Testing 4", "Stand back!")

	// Assert.
	if mock.HelperCallCount != 1 {
		t.Fatalf("expected a HelperCallCount of 1, but was %v", mock.HelperCallCount)
	}

	expectedFatalMessage := "Testing 4\nStand back!\n\nThe provided value was not equal to 5\nExpected: 5\nGot: 4"
	if mock.FatalMessage != expectedFatalMessage {
		t.Fatalf("expected a FatalMessage of:\n%v\nbut returned\n%v", expectedFatalMessage, mock.FatalMessage)
	}
}

// --

type equalToFive struct{}

var _ Assertion = &equalToFive{}

func (*equalToFive) Evaluate(x interface{}) Result {
	if x == 5 {
		return PassedResult
	}

	return Result{
		Pass:    false,
		Message: "The provided value was not equal to 5",
		Specifics: map[string]interface{}{
			"Expected": 5,
			"Got":      x,
		},
	}
}
