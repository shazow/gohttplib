# gohttplib

Shared library that exposes Go's `net/http.Server` with externally-bindable
handlers.

This is a silly project for experimenting with Go v1.5 buildmodes.

**Status**: Tiny subset of the http.HandlerFunc callback gets passed to a C handler callback successfully.


## Getting Started

You'll need Go v1.5 or newer.

```
$ git clone https://github.com/shazow/gohttplib/
$ cd gohttplib
$ make
$ ./build/gohttp-c
```

Now you can request `http://localhost:8000/hello` and the C handler in `examples/c/main.c` will handle it!


## References

* [Go Execution Modes](https://docs.google.com/document/d/1nr-TQHw_er6GOQRsF6T43GGhFDelrAP0NqSS_00RgZQ/view#) document.
* [github.com/jrick/buildmodes](https://github.com/jrick/buildmodes)
* Every single StackOverflow thread tagged [cgo].


## Sponsors

This project was made possible thanks to [Glider Labs](http://gliderlabs.com/).


## License

MIT
