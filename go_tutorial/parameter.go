package main

import "fmt"

// 매개변수 타입, 리턴 타입은 이름 뒤에 지정
func add1(x int, y int) int {
	return x + y
}

// 매개변수 타입 같으면 한 번만 입력
func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add1(42, 13))
	fmt.Println(add2(42, 13))
}
