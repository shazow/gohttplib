from . import route, run

if __name__ == '__main__':
    @route('/')
    def index(w, req):
        w.write("%s %s %s\n" % (req.method, req.host, req.url))
        w.write("Hello, world.\n")

    run()
