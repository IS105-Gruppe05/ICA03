package iso

import "testing"

func TestGreetingExtendedASCIII(t *testing.T) {
	a := isoExtended
	if !(isASCII(a)) {

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
