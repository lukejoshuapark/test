package test

import "testing"

func TestTMockName(t *testing.T) {
	// Arrange.
	mock := NewTMock()

	// Pre-condition.
	name := mock.Name()
	if name != "MockTest" {
		t.Fatalf(`expected default test name to be "MockTest", but was "%v"`, name)
	}

	// Act.
	mock.TestName = "DifferentTest"

	// Assert.
	name = mock.Name()
	if name != "DifferentTest" {
		t.Fatalf(`expected name after change to be "DifferentTest", but was "%v"`, name)
	}
}

func TestTMockHelper(t *testing.T) {
	// Arrange.
	mock := NewTMock()

	// Pre-condition.
	if mock.HelperCallCount != 0 {
		t.Fatalf(`expected initial HelperCallCount value to be 0, but was %v`, mock.HelperCallCount)
	}

	// Act.
	mock.Helper()
	mock.Helper()

	// Assert.
	if mock.HelperCallCount != 2 {
		t.Fatalf(`expected HelperCallCount to be 2 after 2 calls, but was %v`, mock.HelperCallCount)
	}
}

func TestTMockFatalf(t *testing.T) {
	// Arrange.
	mock := NewTMock()

	// Pre-condition.
	if mock.FatalMessage != "" {
		t.Fatalf(`expected initial FatalMessage to be an empty string, but was "%v"`, mock.FatalMessage)
	}

	// Act.
	mock.Fatalf("it blew up %v times", 42)

	// Assert.
	if mock.FatalMessage != "it blew up 42 times" {
		t.Fatalf(`expected FatalMessage to be "it blew up 42 times" but was "%v"`, mock.FatalMessage)
	}
}

func TestTMockFatalfMultipleCalls(t *testing.T) {
	// Arrange.
	mock := NewTMock()
	mock.Fatalf("some arbitrary failure message")

	defer func() {
		r := recover()
		if r != "cannot call Fatalf if previously called" {
			t.Fatalf(`expected the precise panic value "cannot call Fatalf if previously called", but got %v`, r)
		}
	}()

	// Act and Assert.
	mock.Fatalf("another aribitrary failure message")
}
