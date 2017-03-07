package iso

import "testing"

func TestGreetingExtendedASCII(t *testing.T) {
	ascii := GreetingExtendedASCII()
	if !(isASCII(ascii)) {
		t.Fail()
	}
}

func isASCII(s string) bool {
	for _, C := range s {
		if C < 127 {
			return false
		}
	}
	return true
}