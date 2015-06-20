OS  = $(shell uname -s)

ifeq ($(OS), Darwin)
	LDFLAGS += -framework CoreFoundation -framework Security
endif

all: static-lib example-c

build:
	mkdir build

build/gohttplib.a: build/ *.go
	go build -buildmode=c-archive -o build/gohttplib.a

static-lib: build/gohttplib.a

example-c: static-lib examples/c/
	gcc -o build/gohttp-c -I build/ -I . examples/c/main.c build/gohttplib.a $(LDFLAGS) -lpthread

clean:
	rm -rf build/
