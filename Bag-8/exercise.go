package main

import "fmt"

// func main() {
// 	s := "putu gandi mitha wijaya"
// 	for i := range s {
// 		fmt.Printf("%c \n", s[i])
// 		i++
// 	}
// }

func main() {
	for outter := 65; outter <= 90; outter++ {
		for inner := 0; inner <= 2; inner++ {
			fmt.Printf("%#U \n", outter)
		}
		fmt.Println()
	}
}
