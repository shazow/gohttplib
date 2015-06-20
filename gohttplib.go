// TODO: Need to free all C.* types after using them
package main

// #include <stdio.h>
// typedef void FuncPtr(void* w);
// extern void Call_HandleFunc(void* w, FuncPtr* fn);
import "C"
import (
	"net/http"
	"unsafe"
)

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

//export ListenAndServe
func ListenAndServe(caddr *C.char) {
	addr := C.GoString(caddr)
	http.ListenAndServe(addr, nil)
}

//export HandleFunc
func HandleFunc(cpattern *C.char, cfn *C.FuncPtr) {
	pattern := C.GoString(cpattern)
	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		C.Call_HandleFunc(unsafe.Pointer(&responseWriter{w}), cfn)
	})
}

func main() {}
