package main

import "fmt"

// func main() {
// 	if true {
// 		fmt.Println("001")
// 	}
// 	if !false {
// 		fmt.Println("002")
// 	}
// 	if 2 == 2 {
// 		fmt.Println("003")
// 	}
// 	if !(2 == 2) {
// 		fmt.Println("004")
// 	}
// 	if !(2 != 2) {
// 		fmt.Println("005")
// 	}
// }

// deklasari variable didalam if statement
func main() {
	if x := 9; x < 4 {
		fmt.Printf("Benar %d < 4 \n", x)
	} else if x == 7 {
		fmt.Printf("Benar %d == 7 \n", x)
	} else {
		fmt.Printf("Salah karena %d > 4 \n", x)
	}
	//fmt.Println(x) jelas ini akan error karena x didalam block if statement
}
