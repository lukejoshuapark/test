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
		time.Sleep(time.Millisecond * 100)
		atomic.AddInt32(&c, 1)
	}

	f2 := func() {
		time.Sleep(time.Millisecond * 100)
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
	if dt > time.Millisecond*120 || dt < time.Millisecond*80 {
		t.Fatalf("expected a time delta of 100ms but was %v", dt)
	}
}
