package main

import "C"
import "fmt"

//export Print
func Print(s string) {
	fmt.Println(s)
}

func main() {}
