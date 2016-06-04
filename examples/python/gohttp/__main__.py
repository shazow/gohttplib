from . import route, run

if __name__ == '__main__':
    @route('/')
    def index(w, req):
        w.write(b"%s %s %s\n" % (req.method, req.host, req.url))
        w.write(b"Hello, world.\n")

    run()
