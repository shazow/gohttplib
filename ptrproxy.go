package main

import "sync"

func PtrProxy() *ptrProxy {
	return &ptrProxy{
		lookup: map[int]interface{}{},
	}
}

type ptrProxy struct {
	sync.Mutex
	count  int
	lookup map[int]interface{}
}

func (p *ptrProxy) Ref(ptr interface{}) int {
	p.Lock()
	id := p.count
	p.count++
	p.lookup[id] = ptr
	p.Unlock()
	return id
}

func (p *ptrProxy) Deref(id int) (interface{}, bool) {
	p.Lock()
	val, ok := p.lookup[id]
	p.Unlock()
	return val, ok
}

func (p *ptrProxy) Free(id int) {
	p.Lock()
	delete(p.lookup, id)
	p.Unlock()
}
