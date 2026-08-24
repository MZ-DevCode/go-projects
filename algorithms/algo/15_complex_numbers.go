package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	x, y := 3.0, 4.0
	z1 := complex(x, y)
	z2 := 1 + 2i

	sum := z1 + z2
	product := z1 * z2
	modulus := cmplx.Abs(z1)

	re := real(z1)
	im := imag(z1)

	fmt.Println(z1)
	fmt.Println(re, im)
	fmt.Println(sum, product)
	fmt.Println(modulus)
}
