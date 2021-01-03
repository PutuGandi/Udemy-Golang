package main

import (
	"fmt"
)

var bil_bulat_8 int
var bil_cacah_8 uint

func main() {
	bil_bulat_8 = 267 // error karena lebih 127
	bil_bulat_8 = 124 // berhasil karena diantara -128 s/d 127
	fmt.Println(bil_bulat_8)

	bil_cacah_8 = 256 // error karena lebih 255
	bil_cacah_8 = 255 // berhasil karena diantara 0 s/d 255
	fmt.Println(bil_cacah_8)
}
