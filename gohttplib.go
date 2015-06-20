package main

// #include <stdio.h>
// typedef void FuncPtr(void* w);
// extern void Call_HandleFunc(void* w, FuncPtr* fn);
import "C"
import (
	"net/http"
	"unsafe"
)

//export ListenAndServe
func ListenAndServe(caddr *C.char) {
	addr := C.GoString(caddr)
	http.ListenAndServe(addr, nil)
}

//export HandleFunc
func HandleFunc(cpattern *C.char, cfn *C.FuncPtr) {
	pattern := C.GoString(cpattern)
	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		// TODO: Add request to handler API.
		C.Call_HandleFunc(unsafe.Pointer(&responseWriter{w}), cfn)
	})
}

func main() {}
