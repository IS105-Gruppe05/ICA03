package ascii

import "testing"

func TestGreetingASCII(t *testing.T) {
	ascii := a
	if !(isASCII(ascii)) {

		t.Fail()

	}
}

func isASCII(s string) bool {

	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}

/*func test() {
	str := Ascii
	for i := 0; i < len(Ascii); i++ {
		for j := 0; i < len(Ascii); j++ {
			if str[i] != Ascii[j] {
				fmt.Println("fatal")
			}
		}
	}
}
/*

/*func test()
bool Ascii = true
found := Ascii
a := ascii.Ascii
for i  := 0; i < len(Ascii); i++{
	for j := 0; j < len(a); j++ {

		if str[i]
		found = false

		if str[i] != a[j] {
			found = true

			if !found
			fmt.prinf(`%c not found`, str[i])
		}

}
}/*

/*func main() {
a := ascii.Ascii
for i (str)
found = false
for j (a) {

if str[i] == a[j]{
found = true
}
}
if !found {
fmt.Printf(`%c not found`, str[i])
}*/
