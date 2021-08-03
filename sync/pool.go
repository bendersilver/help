package sync

type Pool struct {
	Mutex
}

func (p *Pool) Get() {
	p.Lock()
}

func (p *Pool) Free() {
	p.Unlock()
}
