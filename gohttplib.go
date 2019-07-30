package main

/*
#include <stdlib.h>

typedef struct Request_
{
    const char *Method;
    const char *Host;
		const char *URL;
		const char *Body;
    const char *Headers;
} Request;

typedef unsigned int ResponseWriterPtr;

typedef void FuncPtr(ResponseWriterPtr w, Request *r);

extern void Call_HandleFunc(ResponseWriterPtr w, Request *r, FuncPtr *fn);
*/
import "C"
import (
	"bytes"
	"net/http"
	"unsafe"
	"context"
)

var cpointers = PtrProxy()
var srv http.Server = http.Server{}

//export ListenAndServe
func ListenAndServe(caddr *C.char) {
	addr := C.GoString(caddr)
	srv.Addr = addr
	srv.ListenAndServe()
}

//export Shutdown
func Shutdown() {
	srv.Shutdown(context.Background())
}

//export HandleFunc
func HandleFunc(cpattern *C.char, cfn *C.FuncPtr) {
	// C-friendly wrapping for our http.HandleFunc call.
	pattern := C.GoString(cpattern)
	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		// Convert the headers to a String
		headerBuffer := new(bytes.Buffer)
		req.Header.Write(headerBuffer)
		headersString := headerBuffer.String()
		// Convert the request body to a String
		bodyBuffer := new(bytes.Buffer)
		bodyBuffer.ReadFrom(req.Body)
		bodyString := bodyBuffer.String()
		// Wrap relevant request fields in a C-friendly datastructure.
		creq := C.Request{
			Method:  C.CString(req.Method),
			Host:    C.CString(req.Host),
			URL:     C.CString(req.URL.String()),
			Body:    C.CString(bodyString),
			Headers: C.CString(headersString),
		}
		// Convert the ResponseWriter interface instance to an opaque C integer
		// that we can safely pass along.
		wPtr := cpointers.Ref(unsafe.Pointer(&w))
		// Call our C function pointer using our C shim.
		C.Call_HandleFunc(C.ResponseWriterPtr(wPtr), &creq, cfn)
		// release the C memory
		C.free(unsafe.Pointer(creq.Method))
		C.free(unsafe.Pointer(creq.Host))
		C.free(unsafe.Pointer(creq.URL))
		C.free(unsafe.Pointer(creq.Body))
		C.free(unsafe.Pointer(creq.Headers))
		// Release the ResponseWriter from the registry since we're done with
		// this response.
		cpointers.Free(wPtr)
	})
}

func main() {}
