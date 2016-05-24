OS  = $(shell uname -s)

ifeq ($(OS), Darwin)
	LDFLAGS += -framework CoreFoundation -framework Security
endif

all: examples

examples: example-c example-python

example-c: build/gohttp-c

example-c-static: build/gohttp-c-static

example-python: examples/python/gohttp/_gohttplib.py


build/:
	mkdir build

build/libgohttp.a: *.go
	go build -buildmode=c-archive -o $@

build/libgohttp.so: *.go build
	go build -buildmode=c-shared -o libgohttp.so && mv libgohttp.* $(dir $@)

build/gohttp-c-static: examples/c/ build/libgohttp.a
	gcc -o $@ examples/c/main.c build/libgohttp.a $(LDFLAGS) -lpthread

build/gohttp-c: examples/c/ build/libgohttp.so
	gcc -o $@ examples/c/main.c -Lbuild -lgohttp -lpthread $(LDFLAGS)

examples/python/gohttp/libgohttp.so: build/libgohttp.so
	cp $< $@

examples/python/gohttp/_gohttplib.py: examples/python/gohttp/build_gohttplib.py examples/python/gohttp/libgohttp.so
	cd examples/python && python gohttp/build_gohttplib.py

clean:
	rm -rf build/

benchmark:
	ab -n 10000 -c 10 -s 3 http://127.0.0.1:$(PORT)/
