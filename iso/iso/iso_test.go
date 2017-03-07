package iso

import "testing"

func TestGreetingIso(t *testing.T) {
	iso := a
	if !(isIso(iso)) {

		t.Fail()

	}
}

func isIso(s string) bool {

	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}
