package test

import "fmt"

func main() {
	a := "hello"
	for i := 0; i < len(a); i++ {
		fmt.Printf("%b\n", a[i])
	}
}
