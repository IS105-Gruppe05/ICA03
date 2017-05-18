package ascii

import ("testing"
	"fmt"
)

func TestGreetingASCII(t *testing.T) {

	ascii := ASCIIGreeting
	if !(isASCII(ascii)) {

		t.Fail()

	}
}

func isASCII(s string) bool {

	for _, c := range s {
		if c > 127 {
			fmt.Println("Inneholder ikke bare ASCII tegn\n")
			return false
		}

	}
	fmt.Println("Inneholder bare ASCII tegn\n")
	return true
}
