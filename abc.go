package main

import "fmt"

//Kode som legger på en nedover i ei rekke
/*func main() {
	var extascii []byte
	for i := 0x80; i <= 0x9F; i++ {
		extascii = append(extascii, byte(i))
		fmt.Println(extascii)
	}

}
*/

func main() {
	var extascii []byte
	for i := 0x80; i <= 0xFF; i++ {
		extascii = append(extascii, byte(i))

	}

	//fmt.Println(string(extascii))
	//fmt.Printf("%c", extascii)
	//fmt.println("")
	//fmt.Printf("%s", "\x80")           //printer Euro
	fmt.Printf("%q", string(extascii)) // Genererer ascii tengene "\x00\x01\x02\x03\x04\x05\x06\...

}
