package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x

	for n := 0; n < 10; n++ {
		z = z - ((z*z - x) / (2 * z))
	}
	return z
}

func main() {
	a := 169.
	fmt.Println(Sqrt(a))
	fmt.Println(math.Sqrt(a))
}
