package main

import "fmt"

func main() {
	var point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 6, 7, 9:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}

// func main() {
// 	switch {
// 	case false:
// 		fmt.Println("this should not print")
// 	case (2 == 4):
// 		fmt.Println("this should not print 2")
// 	case (3 == 3):
// 		fmt.Println("prints")
// 		fallthrough
// 	case (4 == 4):
// 		fmt.Println("also true, does it print? ")
// 		fallthrough
// 	case false:
// 		fmt.Println("this is false")
// 	}
// }
