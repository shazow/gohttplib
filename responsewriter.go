package main

// #include <stdio.h>
import "C"
import (
	"net/http"
	"unsafe"
)

//export ResponseWriter_Write
func ResponseWriter_Write(wPtr C.int, cbuf *C.char, length C.int) C.int {
	buf := C.GoBytes(unsafe.Pointer(cbuf), length)

	w, ok := cpointers.Deref(int(wPtr))
	if !ok {
		return C.EOF
	}

	n, err := w.(http.ResponseWriter).Write(buf)
	if err != nil {
		return C.EOF
	}
	return C.int(n)
}

//export ResponseWriter_WriteHeader
func ResponseWriter_WriteHeader(wPtr C.int, header C.int) {
	w, ok := cpointers.Deref(int(wPtr))
	if !ok {
		return
	}
	w.(http.ResponseWriter).WriteHeader(int(header))
}
