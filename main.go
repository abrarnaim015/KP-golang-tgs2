package main

import (
	"fmt"

	"github.com/abrarnaim015/KP-golang-tgs2/helper"
)

func main() {
	num := 17 // Ganti dengan angka yang ingin Anda periksa
	
	a := 5.0
	b := 7.0
	h := 4.0

	fmt.Println(helper.IsPrime(num))
	fmt.Println(helper.Multiple7(num))
	helper.AreaTrapezoid(a, b, h)

}
