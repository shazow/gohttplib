#include "_cgo_export.h"

void Call_HandleFunc(ResponseWriterPtr w, Request *req, FuncPtr *fn) {
    return fn(w, req);
} 
