package main

// typedef char * FuncPtr();
// extern char * Call_HandleFunc(FuncPtr *fn);
import "C"
import (
	"fmt"
	"net/http"
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
		cout := C.Call_HandleFunc(cfn)
		out := C.GoString(cout)
		fmt.Fprintf(w, out)
	})
}

func main() {}
