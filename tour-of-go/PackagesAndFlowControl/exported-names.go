package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(math.pi) 

	// when you export from the package, you should start with capital letter
	// error undefined: math.pi (unexported)

	fmt.Println(math.Pi)

}
