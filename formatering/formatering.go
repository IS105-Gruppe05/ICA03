package main

import "fmt"

func main() {
	/*
		fmt.Printf("%s\n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
		fmt.Printf("%q\n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
		fmt.Printf("%+q\n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
		fmt.Printf("%c\n", []byte("\xc3\xa5"))
		fmt.Printf("%c\n", []byte("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"))
		fmt.Printf("%c", []byte("\xbd"))
		fmt.Printf("%s", []byte("\xb2\x3d\xbc\x20\xe2\x8c\x98"))
	*/
	fmt.Printf("%c %s \n", []byte("\xbd"), "\xb2\x3d\xb2\x20\xe2\x8c\x98")
}
