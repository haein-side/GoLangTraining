package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Println("첫번째")
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // 1. 27 >= 20
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10), // 2. 9
		pow(3, 3, 20), // 3. 20
	)
}
