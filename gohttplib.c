#include "_cgo_export.h"

void Call_HandleFunc(void *w, Request *req, FuncPtr *fn) {
    return fn(w, req);
} 
