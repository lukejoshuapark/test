package test

// Assertion represents a type that is capable of asserting something about a
// value.  A number of built-in assertions are provided.  It is also trivial to
// create new assertions by implementing this interface.
type Assertion interface {
	Evaluate(actual interface{}) Result
}
