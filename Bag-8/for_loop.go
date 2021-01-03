package main

import "fmt"

func main() {
	//outer loop
	// for ini, condidtion, post
	for a := 0; a <= 5; a++ {
		// inner loop
		for b := 0; b <= 3; b++ {
			fmt.Printf("outter loop ialah %d \t inner loop ialah %d \n", a, b)
		}
	}
}
