package main

import "fmt"

func main() {
	a := 5
	b := 10
	fmt.Println("a =", a, "b =", b)
	a, b = b, a
	fmt.Println("a =", a, "b =", b)
}
