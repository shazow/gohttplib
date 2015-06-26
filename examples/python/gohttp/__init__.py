import os
from ._gohttplib import ffi

lib = ffi.dlopen(os.path.join(os.path.dirname(__file__), "libgohttp.so"))


@ffi.callback("void(ResponseWriter*, Request*)")
def handler(w, req):
    body = "Hello, world"
    n = lib.ResponseWriter_Write(w, body, len(body))

    if n < 0:
        print("handler: Failed to write.\n")
        lib.ResponseWriter_WriteHeader(w, 500)
        return

    print("handler: Wrote %d bytes." % n)
