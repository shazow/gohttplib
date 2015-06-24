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

typedef void ResponseWriter;

typedef void FuncPtr(ResponseWriter *w, Request *r);

void ListenAndServe(char* p0);

void HandleFunc(char* p0, FuncPtr* p1);

int ResponseWriter_Write(void* p0, char* p1, int p2);

void ResponseWriter_WriteHeader(void* p0, int p1);
""")


if __name__ == "__main__":
    ffi.compile()
