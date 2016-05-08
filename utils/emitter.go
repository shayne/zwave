package utils

import (
	"time"
)

// Emitter interface
type Emitter interface {
	Emit(next Equatable)
	Reset()
}

type filteredEmitter struct {
	last     Equatable
	lastTime time.Time
	filter   func(Equatable)
}

//
// Creates a new, filtered emitter, such that the wrapped emitter is called at
// most once per minPeriod if the emitted value does not change.

// Filter func
func Filter(emitter func(next Equatable), minPeriod time.Duration) Emitter {
	var f *filteredEmitter

	f = &filteredEmitter{
		last:     nil,
		lastTime: time.Now(),
		filter: func(next Equatable) {
			now := time.Now()
			if f.last != nil &&
				f.last.Equals(next) &&
				now.Sub(f.lastTime) < minPeriod {
			} else {
				f.last = next
				f.lastTime = now
				emitter(next)
			}
		},
	}

	return f

}

func (f *filteredEmitter) Emit(next Equatable) {
	f.filter(next)
}

func (f *filteredEmitter) Reset() {
	f.last = nil
}
