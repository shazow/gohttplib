package main

// #include <stdio.h>
import "C"
import (
	"net/http"
	"unsafe"
)

// responseWriter is an http.ResponseWriter interface container to hold a
// pointer to when passing it into C-land.
// FIXME: Is it possible to go in-out of unsafe.Pointer with an interface
// rather than a literal struct? I couldn't get it to work.
type responseWriter struct {
	http.ResponseWriter
}

//export ResponseWriter_Write
func ResponseWriter_Write(w unsafe.Pointer, cbuf *C.char, length C.int) C.int {
	buf := C.GoBytes(unsafe.Pointer(cbuf), length)
	n, err := (*responseWriter)(w).Write(buf)
	if err != nil {
		return C.EOF
	}
	return C.int(n)
}

//export ResponseWriter_WriteHeader
func ResponseWriter_WriteHeader(w unsafe.Pointer, header C.int) {
	(*responseWriter)(w).WriteHeader(int(header))
}
