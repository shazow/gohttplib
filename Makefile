static-lib:
	go build -buildmode=c-archive

example-c:
	gcc -o gohttp-c examples/c/main.c gohttplib.a -lpthread
