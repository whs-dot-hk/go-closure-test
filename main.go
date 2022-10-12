package main

import "fmt"

type A struct {
	B, C int
}

func main() {
	a := A{1, 1}

	f := func(x int) {
		a.C = x
	}

	fmt.Println(a)

	f(2)

	fmt.Println(a)
}
