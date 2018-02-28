// sem.go -- Counting Semaphores for Go
//
// Licensing Terms: Public Domain
//
// Borrowed from: https://groups.google.com/d/msg/golang-nuts/G5EchCmYy9E/698KiY2TyNgJ
//
// This software does not come with any express or implied
// warranty; it is provided "as is". No claim  is made to its
// suitability for any purpose.

// Package semaphore implements an efficient, simple counting
// semaphore for Go.
package sem

type signal struct{}

// Type representing a semaphore
type Sem struct {
	some  chan signal
	none  chan signal
	count uint32
}

// Create a new counting semaphore instance
func NewSem(n uint32) *Sem {
	s := &Sem{make(chan signal, 1), make(chan signal, 1), n}
	if n == 0 {
		s.none <- signal{}
	} else {
		s.some <- signal{}
	}
	return s
}

// Acquire a semaphore
func (s *Sem) P() {
	<-s.some
	s.count--
	if s.count == 0 {
		s.none <- signal{}
	} else {
		s.some <- signal{}
	}
}

// Release a semaphore
func (s *Sem) V() {
	select {
	case <-s.some:
	case <-s.none:
	}
	s.count++
	s.some <- signal{}
}

// EOF
