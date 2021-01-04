package main

import (
	"fmt"
)

// func main() {
// 	x := "Gandi"
// 	for a, b := range x {
// 		fmt.Printf("abjad %c pada index ke %d desimalnya %v \n", x[a], a, b)
// 		a++
// 	}
// }

func main() {
	x := "putu"
	for a, b := range x {
		fmt.Printf("%v\t%d\t%c\t\t%v\t%d\t%c\n", a, a, a, b, b, b)

	}
}

// func main() {
// 	for i := 33; i <= 122; i++ {
// 		fmt.Printf("Decimal: %v\tHexadecimal: %#x\tUnicode: %#U\n", i, i, i)
// 	}
// }
