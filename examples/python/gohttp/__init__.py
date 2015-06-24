import os
from ._gohttplib import ffi

lib = ffi.dlopen(os.path.join(os.path.dirname(__file__), "libgohttp.so"))
