OS  = $(shell uname -s)

ifeq ($(OS), Darwin)
	LDFLAGS += -framework CoreFoundation -framework Security
endif

all: examples

examples: build/gohttp-c examples/python/gohttp/_gohttplib.so


build/gohttplib.a: *.go
	go build -buildmode=c-archive -o $@

build/libgohttp.so: *.go
	go build -buildmode=c-shared -o $(notdir $@) && mv $(notdir $@) $@

build/gohttp-c: examples/c/ build/gohttplib.a
	gcc -o build/gohttp-c examples/c/main.c build/gohttplib.a $(LDFLAGS) -lpthread


examples/python/gohttp/libgohttp.so: build/libgohttp.so
	cp $< $@

examples/python/gohttp/_gohttplib.so: examples/python/gohttp/*.py examples/python/gohttp/libgohttp.so
	cd $(dir $@)/.. && python gohttp/build_gohttplib.py

clean:
	rm -rf build/
