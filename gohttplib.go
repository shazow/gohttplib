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
	// C-friendly wrapping for our http.HandleFunc call.
	pattern := C.GoString(cpattern)
	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		// Wrap relevant request fields in a C-friendly datastructure.
		creq := C.Request{
			Method: C.CString(req.Method),
			Host:   C.CString(req.Host),
			URL:    C.CString(req.URL.String()),
		}
		// Convert the ResponseWriter interface instance to an opaque C integer
		// that we can safely pass along.
		wPtr := cpointers.Ref(unsafe.Pointer(&w))
		// Call our C function pointer using our C shim.
		C.Call_HandleFunc(C.ResponseWriterPtr(wPtr), &creq, cfn)
		// Release the ResponseWriter from the registry since we're done with
		// this response.
		cpointers.Free(wPtr)
	})
}

func main() {}
