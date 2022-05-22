package test

import (
	"sync"
)

// Concurrently runs all provided functions at the same time in separate
// goroutines.  This call blocks until all the goroutines finish.
func Concurrently(fs ...func()) {
	wg1 := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}
	wg1.Add(1)
	wg2.Add(len(fs))

	for _, f := range fs {
		go func(f func()) {
			defer wg2.Done()
			wg1.Wait()
			f()
		}(f)
	}

	wg1.Done()
	wg2.Wait()
}
