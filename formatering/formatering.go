package main

import "fmt"

func main() {
	fmt.Printf("%c %s \n", []byte("\xbd"), "\xb2\x3d\xb2\x20\xe2\x8c\x98")
}
