package main

import "C"
import "fmt"

//export Hello
func Hello(cs *C.char) {
	s := C.GoString(cs)
	fmt.Println(s)
}

func main() {}
