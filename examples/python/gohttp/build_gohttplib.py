from cffi import FFI

ffi = FFI()
ffi.set_source("gohttp._gohttplib", None)

# Copied from the Go-generated gohttplib.h
ffi.cdef("""
typedef struct Request_
{
    const char *Method;
    const char *Host;
    const char *URL;
} Request;

typedef unsigned int ResponseWriterPtr;

typedef void FuncPtr(ResponseWriterPtr w, Request *r);

void Call_HandleFunc(ResponseWriterPtr w, Request *r, FuncPtr *fn);

void ListenAndServe(char* p0);

void HandleFunc(char* p0, FuncPtr* p1);

int ResponseWriter_Write(unsigned int p0, char* p1, int p2);

void ResponseWriter_WriteHeader(unsigned int p0, int p1);
""")


if __name__ == "__main__":
    ffi.compile()
