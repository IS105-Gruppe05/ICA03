package treasure

import (
	"bytes"
	"fmt"
)

// Kode for Oppgave 3c
// Bruk strengen fra filen treasure.txt som in-data for denne funksjonen
const a = "\x48\x65\x6e\x72\x69\x6b\x20\x41\x72\x6e\x6f\x6c\x64\x20\x57\x65" +
	"\x72\x67\x65\x6c\x61\x6e\x64\x20\x28\x66\xf8\x64\x74\x20\x31\x37\x2e\x20\x6a" +
	"\x75\x6e\x69\x20\x31\x38\x30\x38\x2c\x20\x64\xf8\x64\x20\x31\x32\x2e\x20\x6a" +
	"\x75\x6c\x69\x20\x31\x38\x34\x35\x29\x0a\x56\x69\x20\x65\x72\x65\x20\x65\x6e" +
	"\x20\x6e\x61\x73\x6a\x6f\x6e\x20\x76\x69\x20\x6d\x65\x64\x2c\x0a\x20\x76\x69" +
	"\x20\x73\x6d\xe5\x20\x65\x6e\x20\x61\x6c\x65\x6e\x20\x6c\x61\x6e\x67\x65" +
	"\x2c\x0a\x20\x65\x74\x20\x66\x65\x64\x72\x65\x6c\x61\x6e\x64\x20\x76\x69\x20" +
	"\x66\x72\x79\x64\x65\x73\x20\x76\x65\x64\x2c\x0a\x20\x6f\x67\x20\x76\x69\x2c" +
	"\x20\x76\x69\x20\x65\x72\x65\x20\x6d\x61\x6e\x67\x65\x2e\x0a\x20\x56\xe5\x72" +
	"\x74\x20\x68\x6a\x65\x72\x74\x65\x20\x76\x65\x74\x2c\x20\x76\xe5\x72\x74\x20" +
	"\xf8\x79\x65\x20\x73\x65\x72\x0a\x20\x68\x76\x6f\x72\x20\x67\x6f\x64\x74\x20\x6f" +
	"\x67\x20\x76\x61\x6b\x6b\x65\x72\x74\x20\x4e\x6f\x72\x67\x65\x20\x65\x72\x2c" +
	"\x0a\x20\x76\xe5\x72\x20\x74\x75\x6e\x67\x65\x20\x6b\x61\x6e\x20\x65\x6e\x20" +
	"\x73\x61\x6e\x67\x20\x62\x6c\x61\x6e\x74\x20\x66\x6c\x65\x72\x0a\x20\x61\x76\x20" +
	"\x4e\x6f\x72\x67\x65\x73\x20\xe6\x72\x65\x73\x2d\x73\x61\x6e\x67\x65\x2e\x0a" +
	"\x0a\x20\x4d\x65\x72\x20\x67\x72\xf8\x6e\x74\x20\x65\x72\x20\x67\x72\x65\x73" +
	"\x73\x65\x74\x20\x69\x6e\x67\x65\x6e\x73\x74\x65\x64\x73\x2c\x0a\x20\x6d\x65" +
	"\x72\x20\x66\x75\x6c\x6c\x74\x20\x61\x76\x20\x62\x6c\x6f\x6d\x73\x74\x65\x72" +
	"\x20\x76\x65\x76\x65\x74\x0a\x20\x65\x6e\x6e\x20\x69\x20\x64\x65\x74\x20\x6c" +
	"\x61\x6e\x64\x20\x68\x76\x6f\x72\x20\x6a\x65\x67\x20\x74\x69\x6c\x66\x72\x65" +
	"\x64\x73\x0a\x20\x6d\x65\x64\x20\x66\x61\x72\x20\x6f\x67\x20\x6d\x6f\x72\x20" +
	"\x68\x61\x72\x20\x6c\x65\x76\x65\x74\x2e\x0a\x20\x4a\x65\x67\x20\x76\x69\x6c" +
	"\x20\x64\x65\x74\x20\x65\x6c\x73\x6b\x65\x20\x74\x69\x6c\x20\x6d\x69\x6e\x20" +
	"\x64\xf8\x64\x2c\x0a\x20\x65\x69\x20\x62\x79\x74\x74\x65\x20\x64\x65\x74\x20" +
	"\x68\x76\x6f\x72\x20\x6a\x65\x67\x20\x65\x72\x20\x66\xf8\x64\x64\x2c\x0a\x20" +
	"\x6f\x6d\x20\x6d\x61\x6e\x20\x65\x74\x20\x70\x61\x72\x61\x64\x69\x73\x20\x6d" +
	"\x65\x67\x20\x62\xf8\x64\x0a\x20\x61\x76\x20\x70\x61\x6c\x6d\x65\x72\x20\x6f" +
	"\x76\x65\x72\x73\x76\x65\x76\x65\x74\x2e\x0a"

func PrintTreasureUTF8(treasure_string string) []byte {
	return []byte{'\x00'} // returverdien er her kun en stedsholder
}

func Skrivut() {
	fmt.Printf("%s", a)
	fmt.Printf("%s\n", bytes.Replace([]byte(a), []byte("\xe5"), []byte("\xc3\xa5"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte(a), []byte("\xF8"), []byte("\xc3\x86"), 2))
}

func Replace() {

	s := bytes.Replace([]byte(a), []byte("\xe5"), []byte("\xc3\xa5"), -1)
	e := bytes.Replace([]byte(s), []byte("\xc6"), []byte("\xc3\x86"), -1)
	r := bytes.Replace([]byte(e), []byte("\xf8"), []byte("\xc3\xb8"), -1)
	t := bytes.Replace([]byte(r), []byte("\xd8"), []byte("\xc3\x98"), -1)
	y := bytes.Replace([]byte(t), []byte("\xe6"), []byte("\xc3\xa6"), -1)
	fmt.Printf("%s\n", bytes.Replace([]byte(y), []byte("\xc6"), []byte("\xc3\x86"), -1))

	//func Replace(s, old, new []byte, n int) []byte {
	//bytes.Replace(byteslice, []byte("\xe5"), []byte("\xc3\xa5"), -1)
	//bytes.Replace(byteslice, []byte("\xc6"), []byte("\xc3\x86"), -1)
	//bytes.Replace(byteslice, []byte("\xf8"), []byte("\xc3\xb8"), -1)
	//bytes.Replace(byteslice, []byte("\xd8"), []byte("\xc3\x98"), -1)
	//bytes.Replace(byteslice, []byte("\xe6"), []byte("\xc3\xa6"), -1)
	//bytes.Replace(byteslice, []byte("\xc6"), []byte("\xc3\x86"), -1)
	//}
}
