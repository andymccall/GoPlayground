package main

import "fmt"

func main() {

	// Loop 5 times
	for i := 1;  i<=5; i++ {
		fmt.Printf("Welcome %d times\n",i)
	}

	// Loop forever
	for {
		fmt.Println("I am in love with the Terminal (hit Ctrl+C to stop)")
	}
}
