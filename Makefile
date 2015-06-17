OS  = $(shell uname -s)

ifeq ($(OS), Darwin)
	LDFLAGS += -framework CoreFoundation -framework Security
endif

all: static-lib example-c

static-lib:
	go build -buildmode=c-archive

example-c:
	gcc -o gohttp-c examples/c/main.c gohttplib.a $(LDFLAGS) -lpthread

clean:
	rm *.a *.h gohttp-*
