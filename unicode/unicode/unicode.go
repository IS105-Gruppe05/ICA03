package unicode

import "fmt"

const nor = "\x93\x6e\x6f\x72\x64\x20\x6f\x67\x20\x73\xf8\x72"
const is = "\u0022\x6e\x6f\x72\u00f0\x75\x72\x20\x6f\x67\x20\x73\x75\u00f0\x75\x72\u0022"
const jp = "\u0022\xe5\x8c\x97\xE3\x81\xA8\xE5\x8D\x97\u0022"

//language := "“norður og suður”

func Translate(nor string, language string) string {
	if nor == "nord og sør" {
		if language == "jp" {
			fmt.Printf("%s", jp)
		}
		if language == "is" {
			fmt.Printf("%s", is)
		}

	}
	return ""
}

/*
func Translate(expression string, language string) string {
	if expression == "nord og sør" {
		if language == "is" {
			a := "\x22\x6e\x6f\xf5\x75\x72\x20\x6f\x67\x20\x73\x75\xf5\x75\x73\x22"
			s := bytes.Replace([]byte(a), []byte("\xf5"), []byte("\xc3\xb0"), -1)
			//e := bytes.Replace([]byte(s), []byte("\x93"), []byte(""), -1)
			fmt.Printf("%s\n", s)
		}
	}
	return ""
}
Nord og sør (norsk) "\x93\x6e\x6f\x72\x64\x20\x6f\x67\x20\x73\xf8\x72\x93"
Nord og sør (islands) "\x93\x6e\x6f\xf5\x75\x72\x20\x6f\x20\x73\x75\xf5\x75\x73\x93"
*/
