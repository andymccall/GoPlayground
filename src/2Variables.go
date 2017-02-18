package main

import "fmt"

func main() {
	// String
	var a string = "initial"
	fmt.Println(a)

	// Multiple integers
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Infered type
	var d = true
	fmt.Println(d)

	// Integer, init to 0
	var e int
	fmt.Println(e)

	// Pascal style short declare and assignment
	f := "short"
	fmt.Println(f)
}
