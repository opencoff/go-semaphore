// Test harness for sem.go

package sem_test

import (
	"sync"
	"testing"

	"sem"
)

func Test0(t *testing.T) {
	s := sem.NewSem(0)

	var z int

	var wg sync.WaitGroup

	// Acquire
	go func(s *sem.Sem) {
		wg.Add(1)
		defer wg.Done()

		t.Logf("z=%d\n", z)
		s.P()
		if z != 1 {
			t.Fatalf("release failed\n")
		}
		z = 0
		s.V()
	}(s)

	// Releaser
	go func(s *sem.Sem) {
		wg.Add(1)
		defer wg.Done()
		z = 1
		s.V()
		s.P()
		if z != 0 {
			t.Fatalf("acquire failed\n")
		}
	}(s)

	wg.Wait()
	wg.Wait()

}
