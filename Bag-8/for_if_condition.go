package main

import "fmt"

func main() {
	a := 1
	for a <= 10 {
		if a >= 5 {
			fmt.Printf("ini sudah diatas 5 dan ini looping ke %d \n", a)
			break
		} else {
			fmt.Println(a)
		}
		a++
	}
	fmt.Println("done")
}
