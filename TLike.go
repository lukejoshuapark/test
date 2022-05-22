package test

import "testing"

// TLike defines the methods of *testing.T and *testing.B that are used by this
// package.
//
// Accepting a TLike allows this package to work with both *testing.T
// and *testing.B whilst also enabling internal unit testing with TMock.
type TLike interface {
	Name() string
	Helper()
	Fatalf(format string, args ...interface{})
}

var _ TLike = &testing.T{}
var _ TLike = &testing.B{}
