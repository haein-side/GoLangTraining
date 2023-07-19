package main

import (
	"fmt"
	"math"
)

type Vertex5 struct {
	X, Y float64
}

func Abs(v Vertex5) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex5{3, 4}
	fmt.Println(Abs(v))
}
