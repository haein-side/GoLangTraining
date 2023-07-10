package main

import "fmt"

func main() {
	m4 := make(map[string]int)

	m4["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m4["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m4, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m4["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
