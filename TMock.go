package test

import "fmt"

// TMock is an implementation of TLike that can be inspected for testing
// purposes.
type TMock struct {
	TestName        string
	HelperCallCount int
	FatalMessage    string
}

var _ TLike = &TMock{}

// NewTMock creates a new, empty TMock.
func NewTMock() *TMock {
	return &TMock{
		TestName: "MockTest",
	}
}

// Name returns the value of TestName.  By default, this value is "MockTest".
func (t *TMock) Name() string {
	return t.TestName
}

// Helper increases HelperCallCount by 1.
func (t *TMock) Helper() {
	t.HelperCallCount++
}

// Fatalf stores the error logged in FatalMessage for later inspection.
//
// To reflect the true behavior of Fatalf on a *testing.T or *testing.B, this
// method can only be called once.  Any subsequent calls will panic.
func (t *TMock) Fatalf(format string, args ...interface{}) {
	if t.FatalMessage != "" {
		panic("cannot call Fatalf if previously called")
	}

	t.FatalMessage = fmt.Sprintf(format, args...)
}
