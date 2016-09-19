from . import route, run

if __name__ == '__main__':
    @route('/')
    def index(w, req):
        w.write(b"%s %s %s\n%s\n\n%s\n\n" % (req.method, req.host, req.url, req.headers, req.body))
        w.write(b"Hello, world.\n")

    run()
