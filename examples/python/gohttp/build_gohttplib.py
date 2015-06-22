from cffi import FFI
ffi = FFI()

# Copied from the Go-generated gohttplib.h
lib_header = """
typedef struct Request_
{
    const char *Method;
    const char *Host;
    const char *URL;
} Request;

typedef void ResponseWriter;

typedef void FuncPtr(ResponseWriter *w, Request *r);

extern void ListenAndServe(char* p0);

extern void HandleFunc(char* p0, FuncPtr* p1);

extern int ResponseWriter_Write(void* p0, char* p1, int p2);

extern void ResponseWriter_WriteHeader(void* p0, int p1);
"""
ffi.cdef(lib_header)

ffi.set_source("gohttplib", """
    #include "gohttplib.h"
""", include_dirs=["../../../build"])


if __name__ == "__main__":
    ffi.compile()
