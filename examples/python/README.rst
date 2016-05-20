gohttp: Bindings for gohttplib
==============================

*Warning*: This library currently ships with a shared object of gohttplib
compiled for OSX. It will not work on other platforms at the moment.

See `<https://github.com/shazow/gohttplib>`_ for details.


Usage
-----

::

    from gohttp import route, run
    
    @route('/')
    def index(w, req):
        w.write("%s %s %s\n" % (req.method, req.host, req.url))
        w.write("Hello, world.\n")
    
    run(host='127.0.0.1', port=5000)


License
-------

MIT.
