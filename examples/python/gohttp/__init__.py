import os
from ._gohttplib import ffi

lib = ffi.dlopen(os.path.join(os.path.dirname(__file__), "libgohttp.so"))

# Hang onto handlers so they don't get gc'd
_handlers = []


class ResponseWriter:
    def __init__(self, w):
        self._w = w

    def write(self, body):
        n = lib.ResponseWriter_Write(self._w, body, len(body))
        if n != len(body):
            raise IOError("Failed to write to ResponseWriter.")

    def set_status(self, code):
        lib.ResponseWriter_WriteHeader(self._w, code)


class Request:
    def __init__(self, req):
        self._req = req

    @property
    def method(self):
        return ffi.string(self._req.Method)

    @property
    def host(self):
        return ffi.string(self._req.Host)

    @property
    def url(self):
        return ffi.string(self._req.URL)

    def __repr__(self):
        return "{self.method} {self.url}".format(self=self)


def route(pattern, fn=None):
    """
    Can be used as a decorator.

    :param pattern:
        Address pattern to match against.

    :param fn:
        Handler to call when pattern is matched. Handler is given a
        ResponseWriter and Request object.
    """
    def wrapped(fn):
        @ffi.callback("void(ResponseWriter*, Request*)")
        def handler(w, req):
            fn(ResponseWriter(w), Request(req))

        lib.HandleFunc(pattern, handler)
        _handlers.append(handler)

    if fn:
        return wrapped(fn)

    return wrapped


def run(host='127.0.0.1', port=5000):
    bind = '{}:{}'.format(host or '', port)
    print(" * Running on http://{}/".format(bind))
    lib.ListenAndServe(bind)
