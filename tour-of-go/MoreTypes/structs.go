package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	s := Vertex{}
	s.X = 1
	s.Y = 4

	fmt.Println(s)
}
