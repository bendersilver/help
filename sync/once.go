package sync

import "sync/atomic"

// Once -
type Once struct {
	done uint32
}

// Do -
func (o *Once) Do(f func()) {
	if atomic.AddUint32(&o.done, 1) == 1 {
		f()
	}
}

// OnceQueue -
type OnceQueue struct {
	done uint32
}

// Do -
func (o *OnceQueue) Do(f func()) {
	if atomic.AddUint32(&o.done, 1) == 1 {
		f()
		if atomic.SwapUint32(&o.done, 0) > 1 {
			go o.Do(f)
		}
	}
}
