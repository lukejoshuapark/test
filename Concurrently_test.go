package test

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestConcurrently(t *testing.T) {
	// Arrange.
	c := int32(0)

	// Act.
	f1 := func() {
		time.Sleep(time.Millisecond * 10)
		atomic.AddInt32(&c, 1)
	}

	f2 := func() {
		time.Sleep(time.Millisecond * 10)
		atomic.AddInt32(&c, 1)
	}

	bt := time.Now()
	Concurrently(f1, f2)
	et := time.Now()

	// Assert.
	if c != 2 {
		t.Fatalf("expected c to be 2, but was %v", c)
	}

	dt := et.Sub(bt)
	if dt < time.Millisecond*10 || dt > time.Millisecond*18 {
		t.Fatalf("expected a time delta of 10ms but was %v", dt)
	}
}
