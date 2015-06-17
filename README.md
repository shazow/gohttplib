# gohttplib

Shared library that exposes Go's `net/http.Server` with externally-bindable
handlers.

This is a silly project for experimenting with Go v1.5 buildmodes.

**Status**: Nothing works, just messing around.


## Notes

* Must have a `package main`
* Must have a `func main() {}`
* `go build -buildmode=c-archive` produces `gohttplib.h` boilerplate C headers
  and `gohttplib.a`


## References

* [Go Execution Modes](https://docs.google.com/document/d/1nr-TQHw_er6GOQRsF6T43GGhFDelrAP0NqSS_00RgZQ/view#) document.
* [github.com/jrick/buildmodes](https://github.com/jrick/buildmodes)


## Sponsors

This project was made possible thanks to [Glider Labs](http://gliderlabs.com/).


## License

MIT
