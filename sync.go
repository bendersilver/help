package help

import (
	"sync/atomic"
)

type OnceQueue struct {
	done uint32
}

func (o *OnceQueue) Do(f func()) {
	atomic.AddUint32(&o.done, 1)
	if atomic.LoadUint32(&o.done) == 1 {
		f()
		if atomic.SwapUint32(&o.done, 0) > 1 {
			go o.Do(f)
		}
	}
}
