package main

import (
	"C"
	"sync"
	"unsafe"
)

func PtrProxy() *ptrProxy {
	return &ptrProxy{
		lookup: map[uint]unsafe.Pointer{},
	}
}

type ptrProxy struct {
	sync.Mutex
	count  uint
	lookup map[uint]unsafe.Pointer
}

func (p *ptrProxy) Ref(ptr unsafe.Pointer) C.uint {
	p.Lock()
	id := p.count
	p.count++
	p.lookup[id] = ptr
	p.Unlock()
	return C.uint(id)
}

func (p *ptrProxy) Deref(id C.uint) (unsafe.Pointer, bool) {
	p.Lock()
	val, ok := p.lookup[uint(id)]
	p.Unlock()
	return val, ok
}

func (p *ptrProxy) Free(id C.uint) {
	p.Lock()
	delete(p.lookup, uint(id))
	p.Unlock()
}
