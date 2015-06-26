#!/usr/bin/env python
from gohttp import lib, handler

if __name__ == "__main__":
    lib.HandleFunc("/hello", handler)
    print("Listening on :8001")
    lib.ListenAndServe(":8001")
