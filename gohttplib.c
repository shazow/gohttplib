#include "_cgo_export.h"

void Call_HandleFunc(unsigned int wPtr, Request *req, FuncPtr *fn) {
    return fn(wPtr, req);
} 
