OS  = $(shell uname -s)

ifeq ($(OS), Darwin)
	LDFLAGS += -framework CoreFoundation -framework Security
endif

all: examples

examples: build/gohttp-c examples/python/gohttplib.so


build/gohttplib.a: *.go
	go build -buildmode=c-archive -o build/gohttplib.a

build/gohttp-c: examples/c/ build/gohttplib.a
	gcc -o build/gohttp-c examples/c/main.c build/gohttplib.a $(LDFLAGS) -lpthread

examples/python/gohttplib.so: examples/python/*.py build/gohttplib.a
	cd $(dir $@) && python build_gohttplib.py

clean:
	rm -rf build/
