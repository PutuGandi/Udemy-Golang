package main

import "fmt"

func main() {
	x := 0
	for {
		x++
		if x > 20 {
			break
		}
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("Done")
}
