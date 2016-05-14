package main

import (
	"C"
	"sync"
	"unsafe"
)

// PtrProxy creates a safe pointer registry. It hangs on to an unsafe.Pointer and
// returns a totally-safe C.uint ID that can be used to look up the original
// pointer by using it.
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

// Ref registers the given pointer and returns a corresponding id that can be
// used to retrieve it later.
func (p *ptrProxy) Ref(ptr unsafe.Pointer) C.uint {
	p.Lock()
	id := p.count
	p.count++
	p.lookup[id] = ptr
	p.Unlock()
	return C.uint(id)
}

// Deref takes an id and returns the corresponding pointer if it exists.
func (p *ptrProxy) Deref(id C.uint) (unsafe.Pointer, bool) {
	p.Lock()
	val, ok := p.lookup[uint(id)]
	p.Unlock()
	return val, ok
}

// Free releases a registered pointer by its id.
func (p *ptrProxy) Free(id C.uint) {
	p.Lock()
	delete(p.lookup, uint(id))
	p.Unlock()
}
