package main

import "fmt"

func main() {

	fmt.Printf("%s\n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	//Resultatet
	//er: "??=? ?"

	fmt.Printf("%q\n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	//Resultatet
	//er: "\xbd\xb2=\xbc ?"

	fmt.Printf("%+q\n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
//	"\xbd\xb2=\xbc \u2318"

	fmt.Printf("%c\n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	//Resultatet
	//er: "[� � = � � ]"

	/*fmt.Printf("%s", < bytesekvens >)
	returnerer
	"�?=? ?"*/
}