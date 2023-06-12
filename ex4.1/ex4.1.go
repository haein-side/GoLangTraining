package main

import "fmt"

func main() {
	var a int = 10
	// var : 변수 선언하겠다는 것
	// a : 메모리 공간 가리킴 -> 메모리 공간에 10이 들어가게 됨
	var msg string = "Hello Variable"
	// msg가 가리키는 메모리 공간에 copy 됨
	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)
	// msg가 가리키는 메모리 공간의 데이터 출력됨
}
