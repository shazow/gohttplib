package main

/*
typedef struct Request_
{
    const char *Method;
    const char *Host;
    const char *URL;
} Request;

typedef unsigned int ResponseWriterPtr;

typedef void FuncPtr(ResponseWriterPtr w, Request *r);

extern void Call_HandleFunc(ResponseWriterPtr w, Request *r, FuncPtr *fn);
*/
import "C"
import (
	"net/http"
	"unsafe"
)

var cpointers = PtrProxy()

//export ListenAndServe
func ListenAndServe(caddr *C.char) {
	addr := C.GoString(caddr)
	http.ListenAndServe(addr, nil)
}

//export HandleFunc
func HandleFunc(cpattern *C.char, cfn *C.FuncPtr) {
	pattern := C.GoString(cpattern)
	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		creq := C.Request{
			Method: C.CString(req.Method),
			Host:   C.CString(req.Host),
			URL:    C.CString(req.URL.String()),
		}
		wPtr := cpointers.Ref(unsafe.Pointer(&w))
		C.Call_HandleFunc(C.ResponseWriterPtr(wPtr), &creq, cfn)
		cpointers.Free(wPtr)
	})
}

func main() {}
